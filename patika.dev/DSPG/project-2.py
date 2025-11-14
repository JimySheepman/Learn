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


def numericalize(df, col, name):
    if not is_numeric_dtype(col):
        df[name] = col.cat.codes+1


def numericalize(df, col, name):
    if not is_numeric_dtype(col):
        df[name] = col.cat.codes+1


def proc_df(df, y_fld, nan_dict=None, is_train=True):
    df = df.copy()
    y = df[y_fld].values

    df.drop(y_fld, axis=1, inplace=True)

    if nan_dict is None:
        nan_dict = {}
    for n, c in df.items():
        fix_missing(df, c, n, nan_dict, is_train)
        numericalize(df, c, n)
    if is_train:
        return df, y, nan_dict
    return df, y


def split_train_val(df, n):
    return df[:n].copy(), df[n:].copy()


def apply_cats(df, train):
    for n, c in df.items():
        if train[n].dtype == "category":
            df[n] = pd.Categorical(
                c, categories=train[n].categories, ordered=True)


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


def rmse(x, y):
    return math.sqrt(((x-y)**2).mean())


def get_sample(df, n):
    idxs = np.random.permutation(len(df))[:n]
    return idxs, df.iloc[idxs].copy()


def print_score(m):
    print(f"RMSE of valid set {rmse(m.predict(x_valid),y_valid)}")
    print(f"R^2 of valid set {m.score(x_valid,y_valid)}")
    print(f"RMSE of train set {rmse(m.predict(x_train),y_train)}")
    print(f"R^2 of train set {m.score(x_train,y_train)}")


if __name__ == '__main__':
    df = pd.read_csv("scores.csv", low_memory=False, parse_dates=["age"])
    print(df)
    df = df.sort_values(by="age")
    print(df)
    print(df["gender"].isnull().sum())
    print(df["melanch"].isnull().sum())
    print(df["number"])
    train_cats(df)
    print(df["number"])
    print(df["number"].cat.categories)
    print(df["number"].cat.codes)
    numericalize(df, df["number"], "number")
    print(df["number"])

    for n, c in df.items():
        if is_numeric_dtype(c):
            if df[n].isnull().sum():
                print(n)

    n_valid = 27
    n_train = len(df)-n_valid
    raw_train, raw_valid = split_train_val(df, n_train)
    x_train, y_train, nas = proc_df(raw_train, 'days')
    x_valid, y_valid = proc_df(raw_valid, 'days', nan_dict=nas, is_train=False)
    m = RandomForestRegressor(n_estimators=1, bootstrap=False, n_jobs=-1)
    m.fit(x_train, y_train)
    m.score(x_train, y_train)
    print_score(m)
    m = RandomForestRegressor(n_estimators=10, bootstrap=False, n_jobs=-1)
    m.fit(x_train, y_train)
    print_score(m)
    m = RandomForestRegressor(n_estimators=30, bootstrap=False, n_jobs=-1)
    m.fit(x_train, y_train)
    print_score(m)
    m = RandomForestRegressor(n_estimators=100, bootstrap=False, n_jobs=-1)
    m.fit(x_train, y_train)
    print_score(m)
    idxs, x_train = get_sample(x_train, 3000)
    y_train = y_train[idxs]
    m = RandomForestRegressor(n_estimators=700, n_jobs=-1)
    m.fit(x_train, y_train)
    print_score(m)
