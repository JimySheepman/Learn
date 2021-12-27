import math
import numpy as np
import pandas as pd
from sklearn import metrics
from IPython.display import display
from sklearn.metrics import r2_score
from pandas.api.types import is_string_dtype
from pandas.api.types import is_numeric_dtype
from sklearn.ensemble import RandomForestRegressor


def rmse(x, y):
    return math.sqrt(((x-y)**2).mean())


def print_score(m):
    print(f"RMSE of valid {rmse(m.predict(df.drop('Y2',axis=1)),df.Y2)}")
    print(f"R^2 of valid {m.score(df.drop('Y2',axis=1),df.Y2)}")


if __name__ == '__main__':
    df = pd.read_csv("ENB2012_data.csv")
    print(df)
    print(df.isnull().sum())
    print(df.Y2.mean())
    print(df.Y2-df.Y2.mean())
    print((df.Y2-df.Y2.mean()).pow(2).sum())
    filt = (df.X5 < 4)
    df1 = df[filt]
    df2 = df[~filt]
    print(df1)
    print(df2)
    print(df1.Y2.mean())
    print(df2.Y2.mean())
    print((df1.Y2-df1.Y2.mean()).pow(2).sum() +
          (df2.Y2-df2.Y2.mean()).pow(2).sum())
    filt1 = (df1.X8 <= 3)
    df2 = df1[filt1]
    df3 = df1[~filt1]
    print(df2)
    print(df3)
    print((df1.Y2-df1.Y2.mean()).pow(2).sum()+(df2.Y2-df2.Y2.mean()
                                               ).pow(2).sum()+(df3.Y2-df3.Y2.mean()).pow(2).sum())
    filt2 = (df3.X7 <= 0.25)
    df4 = df3[filt2]
    df5 = df3[~filt2]
    print(df4.Y2.mean())
    print(df5.Y2.mean())
    print((df1.Y2-df1.Y2.mean()).pow(2).sum()+(df2.Y2-df2.Y2.mean()).pow(2).sum() +
          (df4.Y2-df4.Y2.mean()).pow(2).sum()+(df5.Y2-df5.Y2.mean()).pow(2).sum())
    df["Y2"] = np.log(df.Y2)
    m = RandomForestRegressor(n_estimators=26, bootstrap=True, n_jobs=-1)
    m.fit(df.drop('Y2', axis=1), df.Y2)
    print(m.score(df.drop('Y2', axis=1), df.Y2))
    print(rmse(m.predict(df.drop('Y2', axis=1)), df.Y2))
    print_score(m)
