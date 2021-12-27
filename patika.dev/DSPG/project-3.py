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


def apply_cats(df, train):
    for n, c in df.items():
        if train[n].dtype == "category":
            df[n] = pd.Categorical(
                c, categories=tran[n].categories, ordered=True)


def numericalize(df, col, name):
    if not is_numeric_dtype(col):
        df[name] = col.cat.codes+1


def numericalize(df, col, name):
    if not is_numeric_dtype(col):
        df[name] = col.cat.codes+1


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


def rmse(x, y):
    return math.sqrt(((x-y)**2).mean())


def print_score(m):
    print(f"RMSE of valid set {rmse(m.predict(x_valid),y_valid)}")
    print(f"R^2 of valid set {m.score(x_valid,y_valid)}")
    print(f"RMSE of train set {rmse(m.predict(x_train),y_train)}")
    print(f"R^2 of train set {m.score(x_train,y_train)}")


if __name__ == '__main__':
    df = pd.read_csv("Train.csv", low_memory=False)
    print(df)
    print(df.isnull().sum())
    train_cats(df)
    print(df["image_id"])
    print(df["image_id"].cat.codes)
    numericalize(df, df["image_id"], "image_id")
    print(df["image_id"])
    df = df.sort_values(by="image_id")
    print(df)
    n_valid = 100
    n_train = len(df)-n_valid
    raw_train, raw_valid = split_train_val(df, n_train)
    x_train, y_train, nas = proc_df(raw_train, 'healthy')
    x_valid, y_valid = proc_df(
        raw_valid, 'healthy', nan_dict=nas, is_train=False)
    m = RandomForestRegressor(n_estimators=1000, n_jobs=-1)
    m.fit(x_train, y_train)
    m.score(x_train, y_train)
    print_score(m)
    print(x_train)
    print(y_train)
    print(x_valid)
    print(y_valid)
    print(df)
    print(df["healthy"].mean())
    print(df["healthy"]-df["healthy"].mean())
    print(math.sqrt(((df.healthy-df.healthy.mean()).pow(2).sum())/len(df)))
    filt = (df.rust > 0)
    df2 = df[filt]
    df3 = df[~filt]
    print(df2)
    print(math.sqrt(((df2.healthy-df2.healthy.mean()).pow(2).sum() +
          (df3.healthy-df3.healthy.mean()).pow(2).sum())/len(df)))
    filt2 = (df3.scab > 0)
    df4 = df3[filt2]
    df5 = df3[~filt2]
    print(df4)
    print(df5)
    print(math.sqrt(((df2.healthy-df2.healthy.mean()).pow(2).sum()+(df4.healthy -
          df4.healthy.mean()).pow(2).sum()+(df5.healthy-df5.healthy.mean()).pow(2).sum())/len(df)))
