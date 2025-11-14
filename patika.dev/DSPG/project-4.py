import math
import numpy as np
import pandas as pd
from sklearn import metrics
from IPython.display import display
from pandas.api.types import is_string_dtype
from pandas.api.types import is_numeric_dtype
from sklearn.ensemble import RandomForestRegressor


def train_cats(df):
    for n, c in df.items():
        if is_string_dtype(c):
            df[n] = c.astype("category").cat.as_ordered()


def fix_missing(df, col, name, nan_dict, is_train):
    if is_numeric_dtype(col):
        if pd.isnull(col).sum():
            df[name+"_na"] = pd.isnull(col)
            nan_dict[name] = col.median()
            df[name] = col.fillna(nan_dict[name])
    else:
        if is_numeric_dtype(col):
            if name in nan_dict:
                df[name+"_na"] = pd.isnull(col)
                df[name] = col.fillna(nan_dict[name])
            else:
                df[name] = col.fillna(df[name].median())


def numericalize(df, col, name):
    if not is_numeric_dtype(col):
        df[name] = col.cat.codes+1


def split_train_val(df, n):
    return df[:n].copy(), df[n:].copy()


def rmse(x, y):
    return math.sqrt(((x-y)**2).mean())


def print_score(m):
    print(f"RMSE of valid set {rmse(m.predict(x_valid),y_valid)}")
    print(f"R^2 of valid set {m.score(x_valid,y_valid)}")
    print(f"RMSE of train set {rmse(m.predict(x_train),y_train)}")
    print(f"R^2 of train set {m.score(x_train,y_train)}")


if __name__ == '__main__':

    df2 = pd.read_csv("energy_dataset.csv")
    df1 = pd.read_csv("weather_features.csv")
    df = pd.concat([df1, df2], axis=1, join='inner')
    print(df)
    print(df.isnull().sum())
    train_cats(df)
    df["city_name"].cat.codes
    df["dt_iso"].cat.codes
    df["city_name"].cat.categories
    df3 = df[["dt_iso", "city_name", "temp", "temp_min",
              "temp_max", "humidity", "price day ahead"]]
    df3["city_name"].cat.categories
    df3["city_name"].cat.codes
    df = df3
    df = df = df.sort_values(by="dt_iso", ascending=False)
    print(df)
    n_train = 16200
    n_valid = len(df)-n_train
    raw_train, raw_valid = split_train_val(df, n_train)
    x_train, y_train, nas = proc_df(raw_train, 'price day ahead')
    x_valid, y_valid = proc_df(
        raw_valid, 'price day ahead', nan_dict=nas, is_train=False)
    x_valid
    y_valid
    print(df["price day ahead"].mean().sum())
    print(df["price day ahead"].mean().sum()-df["price day ahead"])
    df = df.rename({'price day ahead': 'price'}, axis=1)
    print(math.sqrt(((df.price-df.price.mean()).pow(2).sum())/len(df)))
    df["Temp_Cel"] = df["temp"].values-270
    print(df["Temp_Cel"].mean().sum())
    filt = (df.Temp_Cel > 21)
    df1 = df[filt]
    df2 = df[~filt]
    print(math.sqrt(((df.price-df.price.mean()).pow(2).sum())/len(df)))
    print(math.sqrt(((df1.price-df1.price.mean()).pow(2).sum() +
          (df2.price-df2.price.mean()).pow(2).sum())/len(df)))
    filt1 = (df2.Temp_Cel < 15)
    df3 = df2[filt1]
    df4 = df2[~filt1]
    print(df3)
    print(df4)
    print(math.sqrt(((df1.price-df1.price.mean()).pow(2).sum()+(df3.price -
          df3.price.mean()).pow(2).sum()+(df4.price-df4.price.mean()).pow(2).sum())/len(df)))
    filt2 = (df1.Temp_Cel < 27)
    df5 = df1[filt2]
    df6 = df1[~filt2]
    print(df5)
    print(df6)
    print(math.sqrt(((df5.price-df5.price.mean()).pow(2).sum()+(df6.price-df6.price.mean()).pow(2).sum() +
          (df3.price-df3.price.mean()).pow(2).sum()+(df4.price-df4.price.mean()).pow(2).sum())/len(df)))
    m = RandomForestRegressor(n_estimators=1, bootstrap=False, n_jobs=-1)
    m.fit(x_train, y_train)
    m.score(x_train, y_train)
    print_score(m)
    m = RandomForestRegressor(n_estimators=10, bootstrap=False, n_jobs=-1)
    m.fit(x_train, y_train)
    print_score(m)
    m = RandomForestRegressor(n_estimators=100, bootstrap=False, n_jobs=-1)
    m.fit(x_train, y_train)
    print_score(m)
    m = RandomForestRegressor(n_estimators=1000, bootstrap=False, n_jobs=-1)
    m.fit(x_train, y_train)
    print_score(m)
