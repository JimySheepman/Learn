import const
from util.urls import URLUtility

util = URLUtility(const.ApodEclipseImage())
print(len(util.data))
