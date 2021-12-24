from sklearn.linear_model import LinearRegression
import warnings
import pandas as pd
import numpy as np
import matplotlib.pyplot as plt
plt.style.use("fivethirtyeight")

warnings.filterwarnings("ignore")


df1 = pd.read_csv("2015.csv")
df2 = pd.read_csv("2016.csv")
df3 = pd.read_csv("2017.csv")
df4 = pd.read_csv("2018.csv")
df5 = pd.read_csv("2019.csv")

print(df1.columns, "\n", df2.columns, "\n",
      df3.columns, "\n", df4.columns, "\n", df5.columns)

df1.drop(["Region"], axis=1, inplace=True)
df1["Year"] = 2015
df1.rename(columns={"Happiness Rank": "Rank",
                    "Happiness Score": "Score",
                    "Standard Error": "Error",
                    "Family": "social_support",
                    "Economy (GDP per Capita)": "Economy",
                    "Health (Life Expectancy)": "life_expectancy",
                    "Trust (Government Corruption)": "Corruption",
                    "Dystopia Residual": "dsytopia_residual"},
           inplace=True)

df2["Error"] = df2["Upper Confidence Interval"] - \
    df2["Lower Confidence Interval"]
df2.drop(["Region", "Upper Confidence Interval",
         "Lower Confidence Interval"], axis=1, inplace=True)
df2["Year"] = 2016
df2.rename(columns={"Happiness Rank": "Rank",
                    "Happiness Score": "Score",
                    "Family": "social_support",
                    "Economy (GDP per Capita)": "Economy",
                    "Health (Life Expectancy)": "life_expectancy",
                    "Trust (Government Corruption)": "Corruption",
                    "Dystopia Residual": "dsytopia_residual"},
           inplace=True)

df3["Error"] = df3["Whisker.high"] - df3["Whisker.low"]
df3.drop(["Whisker.high", "Whisker.low"], axis=1, inplace=True)
df3["Year"] = 2017
df3.rename(columns={"Happiness.Rank": "Rank",
                    "Happiness.Score": "Score",
                    "Family": "social_support",
                    "Economy..GDP.per.Capita.": "Economy",
                    "Health..Life.Expectancy.": "life_expectancy",
                    "Trust..Government.Corruption.": "Corruption",
                    "Dystopia.Residual": "dsytopia_residual"},
           inplace=True)

df4["Year"] = 2018
df4.rename(columns={"Overall rank": "Rank",
                    "Country or region": "Country",
                    "GDP per capita": "Economy",
                    "Social support": "social_support",
                    "Healthy life expectancy": "life_expectancy",
                    "Freedom to make life choices": "Freedom",
                    "Perceptions of corruption": "Corruption"},
           inplace=True)

df5["Year"] = 2019
df5.rename(columns={"Overall rank": "Rank",
                    "Country or region": "Country",
                    "GDP per capita": "Economy",
                    "Social support": "social_support",
                    "Healthy life expectancy": "life_expectancy",
                    "Freedom to make life choices": "Freedom",
                    "Perceptions of corruption": "Corruption"},
           inplace=True)

cols = ["Rank", "Country", "Economy", "social_support",
        "life_expectancy", "Freedom", "Generosity", "Corruption", "Year"]
data = pd.concat([df1[cols], df2[cols], df3[cols], df4[cols],
                 df5[cols]], axis=0).reset_index(drop=True)
print(data.sample(10))
print(data.isnull().sum())


d = data[data["Country"] == "United Arab Emirates"][["Corruption", "Year"]]

plt.figure(figsize=(10, 7))
plt.plot(d["Year"], d["Corruption"], "*", color="teal")
plt.axis([2014, 2020, 0, 0.5])
plt.title("Corruption - United Arab Emirates")
plt.show()


X = d[d["Year"] != 2018]["Year"].values.reshape(-1, 1)
y = d[d["Year"] != 2018]["Corruption"].values

reg = LinearRegression().fit(X, y)
print("2018 için tahmin: ", reg.predict(np.array(2018).reshape(-1, 1)))

i = data[data["Corruption"].isnull()].index

data["Corruption"][i] = reg.predict(np.array(2018).reshape(-1, 1))
data[data["Country"] == "United Arab Emirates"]

print("Toplam null girdi sayısı:", data.isnull().sum().sum())


grouped_country = data.groupby("Country").mean()

fig = plt.figure(figsize=(15, 20))

rows = 3
columns = 2

for i, column in enumerate(grouped_country.columns.values[1:-1]):
    if column != "Corruption":
        sorted_grouped_country = grouped_country.sort_values(
            column, ascending=False)[:10]
    else:
        sorted_grouped_country = grouped_country.sort_values(column)[:10]
    plt.subplot(rows, columns, i+1)
    sorted_grouped_country[column].plot(kind="barh", color="indianred")
    plt.ylabel("")
    plt.title(column)

    fig = plt.figure(figsize=(15, 20))

rows = 3
columns = 2

for i, column in enumerate(grouped_country.columns.values[1:-1]):
    if column != "Corruption":
        sorted_grouped_country = grouped_country.sort_values(
            column, ascending=False).tail(10)
    else:
        sorted_grouped_country = grouped_country.sort_values(column).tail(10)
    plt.subplot(rows, columns, i+1)
    sorted_grouped_country[column].plot(kind="barh", color="mediumseagreen")
    plt.ylabel("")
    plt.title(column)


grouped_years = data.sort_values(["Year", "Rank"]).groupby(
    "Year").head(10).reset_index(drop=True)
print(grouped_years)


colors = ["palevioletred", "deepskyblue", "mediumseagreen", "steelblue",
          "goldenrod", "mediumorchid", "lightseagreen", "cornflowerblue", "salmon", "teal"]

fig = plt.figure(figsize=(15, 20))

for i, country in enumerate(grouped_years["Country"].unique()[:10]):
    plt.subplot(5, 2, i+1)
    plt.scatter(grouped_years[grouped_years["Country"] == country]["Year"], grouped_years[grouped_years["Country"]
                == country]["Rank"], s=100, marker="*", linestyle="dotted", c=colors[i])
    plt.plot(grouped_years[grouped_years["Country"] == country]["Year"],
             grouped_years[grouped_years["Country"] == country]["Rank"], linestyle="dotted", c=colors[i])
    plt.axis([2014, 2019, 0, 11])
    plt.title(country)
    plt.suptitle("Top 10 country rankings through years")
    plt.tight_layout()


grouped_years_ = data.sort_values(["Year", "Rank"]).groupby(
    "Year").tail(10).reset_index(drop=True)
print(grouped_years_)


fig = plt.figure(figsize=(15, 20))

for i, country in zip(range(10), grouped_years_["Country"].unique()):
    plt.subplot(5, 2, i+1)
    plt.scatter(grouped_years_[grouped_years_["Country"] == country]["Year"], grouped_years_[
                grouped_years_["Country"] == country]["Rank"], s=100, marker="*", linestyle="dotted", c=colors[i])
    plt.plot(grouped_years_[grouped_years_["Country"] == country]["Year"], grouped_years_[
             grouped_years_["Country"] == country]["Rank"], linestyle="dotted", c=colors[i])
    plt.axis([2014, 2019, 145, 160])
    plt.title(country)
    plt.suptitle("Country rankings having lowest scores through years")
    plt.tight_layout()
