import pandas as pd

planets_df =pd.read_csv("http://localhost:8081/planets_pandas.csv",index_col='Name')
print(planets_df)