import warnings
import seaborn as sns
import pandas as pd
import numpy as np
import matplotlib.pyplot as plt
plt.style.use("fivethirtyeight")
warnings.filterwarnings('ignore')


df1 = pd.read_csv("accidents_2009_to_2011.csv\\accidents_2009_to_2011.csv")
df2 = pd.read_csv("accidents_2012_to_2014.csv\\accidents_2012_to_2014.csv")
all_set = pd.concat([df1, df2], axis=0)
data = all_set.copy()
print("----------------------------------<2009-2011>----------------------------------------", "\n")
print(df1.info(), "\n")
print(f"Total null entries at df1: {df1.isna().sum().sum()} \n")
print("----------------------------------<2012-2014>----------------------------------------", "\n")
print(df2.info(), "\n")
print(f"Total null entries at df2: {df2.isna().sum().sum()} \n")
for i in data.columns:
    print("{} ({}) : \n{}".format(i, len(data[i].unique()), data[i].unique()))
    print("-----------------------------------------------------------------")


data.dropna(axis=1, how="all", inplace=True)
data.drop(["Accident_Index", "LSOA_of_Accident_Location"],
          axis=1, inplace=True)
data["Date"] = pd.to_datetime(data["Date"])
data["Month"] = data["Date"].dt.month
data["Day"] = data["Date"].dt.day
data["Hour"] = pd.to_datetime(data["Time"]).dt.hour
data["Hour"] = pd.cut(data["Hour"], bins=[0, 4, 8, 12, 16, 20, 24], include_lowest=True, labels=[
                      "0-4", "4-8", "8-12", "12-16", "16-20", "20-24"])
data.drop("Date", axis=1, inplace=True)
print(data.head())


fig = plt.figure(figsize=(16,12))

plt.subplot(2,2,1)
data["Police_Force"].value_counts().plot(kind="bar", color="mediumseagreen")
plt.title("Police Force")

plt.subplot(2,2,2)
data["Accident_Severity"].value_counts().plot(kind="bar", color="mediumseagreen")
plt.title("Accident Severity")

plt.subplot(2,2,3)
data["Number_of_Vehicles"].value_counts().plot(kind="bar", color="mediumseagreen")
plt.title("Number of Vehicles")

plt.subplot(2,2,4)
data["Number_of_Casualties"].value_counts().plot(kind="bar", color="mediumseagreen")
plt.title("Number of Casualties")



print("Boylam aralığı: ",min(data["Longitude"]), ",", max(data["Longitude"]))
print("Enlem aralığı: ",min(data["Latitude"]), ",", max(data["Latitude"]))



data["Longitude_"] = pd.cut(data["Longitude"], 8, include_lowest=True, ordered=True)
data["Latitude_"] = pd.cut(data["Latitude"], 10, include_lowest=True, ordered=True)
print(data.head())



print("Kazaların en çok gerçekleştiği bölgenin coğrafi koordinatları:\n")
print(data["Longitude_"].value_counts(),"\n")
print(data["Latitude_"].value_counts())



fig = plt.figure(figsize=(25,17))

plt.subplot(2,2,1)
data["Weather_Conditions"].value_counts().plot(kind="barh", color="deepskyblue")
plt.title("Weather Conditions")

plt.subplot(2,2,2)
data["Light_Conditions"].value_counts().plot(kind="barh", color="deepskyblue")
plt.title("Light Conditions")

plt.subplot(2,2,3)
data["Road_Surface_Conditions"].value_counts().plot(kind="barh", color="deepskyblue")
plt.title("Road Surface Conditions")

plt.subplot(2,2,4)
data["Special_Conditions_at_Site"].value_counts().plot(kind="barh", color="deepskyblue")
plt.title("Special Conditions")

plt.show()



fig = plt.figure(figsize=(25,17))

plt.subplot(2,2,1)
data["Road_Type"].value_counts().plot(kind="barh", color="palevioletred")
plt.title("Road Type")

plt.subplot(2,2,2)
data["Speed_limit"].value_counts().plot(kind="barh", color="palevioletred")
plt.title("Speed Limit")

plt.subplot(2,2,3)
data["Junction_Control"].value_counts().plot(kind="barh", color="palevioletred")
plt.title("Junction Control")

plt.subplot(2,2,4)
data["Pedestrian_Crossing-Human_Control"].value_counts().plot(kind="barh", color="palevioletred")
plt.title("Pedestrian Crossing - Human Control")

plt.show()

urban = data[data["Urban_or_Rural_Area"] == 1]
rural = data[data["Urban_or_Rural_Area"] == 2]

plt.figure(figsize=(15,9))
plt.plot(data["Year"].unique(), urban["Year"].value_counts(sort=False), "-.", label="Urban", color="darkorange")
plt.plot(data["Year"].unique(), rural["Year"].value_counts(sort=False), "-.", label="Rural", color="teal")
plt.legend(loc="best")
plt.title("Accident Counts")
plt.show()




data["Day_of_Week"] = data["Day_of_Week"].astype(str)
data["Day_of_Week"] = data["Day_of_Week"].map({"1": "Monday", "2": "Tuesday", "3": "Wednesday", "4": "Thursday", "5": "Friday", "6": "Saturday", "7":"Sunday"})

data["Month"] = data["Month"].astype(str)
data["Month"] = data["Month"].map({"1": "January", "2": "February", "3": "March", "4": "April", "5": "May", "6": "June", "7": "July", "8": "August", "9": "September", "10": "October", "11": "November", "12": "December"})

fig = plt.figure(figsize=(25,17))

plt.subplot(2,2,1)
data["Year"].value_counts().plot(kind="barh", color="indianred")
plt.title("Years and Accident Counts")

plt.subplot(2,2,2)
data["Month"].value_counts().plot(kind="barh", color="indianred")
plt.title("Months and Accident Counts")

plt.subplot(2,2,3)
data["Day_of_Week"].value_counts().plot(kind="barh", color="indianred")
plt.title("Days and Accident Counts")

plt.subplot(2,2,4)
data["Hour"].value_counts().plot(kind="barh", color="indianred", sharex=False)
plt.title("Hours and Accident Counts")

plt.show()



colors = ["deepskyblue", "sandybrown", "yellowgreen", "cornflowerblue", "mediumpurple", "salmon", "mediumorchid"]

for i,day in enumerate(data["Day_of_Week"].unique()):
    plt.figure(figsize=(12,7))
    data[data["Day_of_Week"] == day]["Time"].value_counts(sort=False).plot(color=colors[i])
    plt.xlabel(day)
    plt.show()

