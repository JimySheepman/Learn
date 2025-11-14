# %% [markdown]
# # Fundamantels

# %%
import sqlite3
import json
from operator import methodcaller
from operator import mul
from operator import add
from operator import itemgetter
from itertools import groupby
from collections import deque
import collections
from math import pow
import cmath
import sys
from math import ceil
from math import sqrt       # Recommended
from math import sqrt, ceil  # Not recommended
from math import *
from math import factorial as fact, gamma, atan as arctan
from urllib.request import urlopen as geturl, pathname2url as path2url, getproxies
from hello import *
from urllib.request import urlopen
from math import sin
import numpy as np
import random
import string
from types import MethodType
from functools import reduce
import os
from array import *
import math
from enum import Enum
import pytz
from dateutil.parser import parse
from datetime import date
import calendar
import time
from dateutil import tz
from datetime import datetime
from datetime import datetime, timedelta, timezone
import dateutil.parser
import datetime
import hello as say
import keyword


# %%
a = 2
print(a)

# %%
b = 1000000000000123123
b

# %%
pi = 3.14
pi

# %%
c = 'A'
c

# %%
name = 'merlins'
name

# %%
q = True
x = None
q, x

# %%
print(keyword.kwlist)

# %%
a = b = c = 1
print(a, b, c)

# %%
a, b, _ = 1, 2, 3
print(a, b)


# %%
x = [1, 2, 3]
x

# %%
x = [1, 2, [1, 2, 3], 3]
x

# %%
x[2], x[2][2]

# %%
x = True
y = False

x or y, x and y, not x

# %%
a = 2.0
b = 100.e0
c = 123456789.e1
a, b, c

# %%
a = 2 + 1j
b = 100 + 10j
a, b

# %%
a = (1, 2, 3)
b = ('a', 1, 'python', (1, 2))
a, b

# %%
a = [1, 2, 3]
b = ['a', 1, 'python', (1, 2), [1, 2]]
b[2] = 'something else'  # allowed
a, b

# %%
a = {1, 2, 'a'}
a

# %%
a = {1: 'one',
     2: 'two'}
b = {'a': [1, 2, 3],
     'b': 'a string'}
a, b

# %%
type(a)

# %%
i = 7
if isinstance(i, int):
    i += 1
elif isinstance(i, str):
    i = int(i)
    i += 1

i

# %%
x = None
if x is None:
    print('Not a surprise, I just defined x as None.')

# %%
a = 'hello'
print(a, type(a))
print(list(a), type(list(a)))
print(set(a), type(set(a)))
print(tuple(a), type(tuple(a)))

# %%
print(dir(list))

# %%
print(dir(dict))

# %%
print(dir(tuple))

# %%
print(dir(set))

# %%
name = input("Enter the name: ")
name

# %%
x = int(input("Enter the number: "))
x = x**2
print("Square number: ", x)

# %%
print(dir(__builtins__))

# %%
help(zip)

# %%
help(set)

# %%

say.hello_world()

# %%
s = """w'o"w"""
repr(s), str(s)

# %%
today = datetime.datetime.now()
str(today), repr(today)

# %%
a_str = 'hello world'
a_str, a_str[2:5]

# %%
basket = {'apple', 'orange', 'apple', 'pear', 'orange', 'banana'}
basket

# %%
a = set('helloWorld')
a

# %%
a.add('t')
a

# %%
b = frozenset('asdfagsa')
b

# %%
int_number = 10
float_number = 10.4
complex_number = 10.4j
long_number = 2**61 - 1

# %%
_list = [123, 'abcd', 10.2, 'd']
list1 = ['hello', 'world']
_list+list1

# %%
dic = {'name': 'red', 'age': 10}
dic.values(), dic.keys()

# %%
tuple1 = ((123, 'hello'), 'world')
tuple1[0][0], tuple1[0]


# %%
def func():
    """This is a function that does nothing at all"""
    return


print(func.__doc__)

# %%
help(func)

# %%


def greet(name, greeting="Hello"):
    """Print a greeting to the user `name`
    Optional parameter `greeting` can change what they're greeted with."""
    print("{} {}".format(greeting, name))


help(greet)

# %%


def greet(name, greeting="Hello"):
    # Print a greeting to the user `name`
    # Optional parameter `greeting` can change what they're greeted with.
    print("{} {}".format(greeting, name))


print(greet.__doc__)


# %% [markdown]
# ## Date and Time

# %%
dt = datetime.datetime.strptime(
    "2016-04-15T08:27:18-0500", "%Y-%m-%dT%H:%M:%S%z")
dt

# %%
dt = dateutil.parser.parse("2016-04-15T08:27:18-0500")
dt

# %%

JST = timezone(timedelta(hours=+9))
dt = datetime(2015, 1, 1, 12, 0, 0, tzinfo=JST)
print(dt)  # 2015-01-01 12:00:00+09:00
print(dt.tzname())  # UTC+09:00
dt = datetime(2015, 1, 1, 12, 0, 0, tzinfo=timezone(timedelta(hours=9), 'JST'))
print(dt.tzname)  # 'JST'

# %%

local = tz.gettz()  # Local time
PT = tz.gettz('US/Pacific')  # Pacific time
dt_l = datetime(2015, 1, 1, 12, tzinfo=local)  # I am in EST
dt_pst = datetime(2015, 1, 1, 12, tzinfo=PT)
dt_pdt = datetime(2015, 7, 1, 12, tzinfo=PT)  # DST is handled automatically
print(dt_l)
print(dt_pst)
print(dt_pdt)


# %%

utc = tz.tzutc()
local = tz.tzlocal()
utc_now = datetime.utcnow()
utc_now  # Not timezone-aware.
utc_now = utc_now.replace(tzinfo=utc)
utc_now  # Timezone-aware.
local_now = utc_now.astimezone(local)
local_now  # Converted to local time.


# %%

today = datetime.date.today()
print('Today:', today)
yesterday = today - datetime.timedelta(days=1)
print('Yesterday:', yesterday)
tomorrow = today + datetime.timedelta(days=1)
print('Tomorrow:', tomorrow)
print('Time between tomorrow and yesterday:', tomorrow - yesterday)


# %%

seconds_since_epoch = time.time()
utc_date = datetime.utcfromtimestamp(seconds_since_epoch)
utc_date


# %%


def monthdelta(date, delta):
    m, y = (date.month+delta) % 12, date.year + ((date.month)+delta-1) // 12
    if not m:
        m = 12
    d = min(date.day, calendar.monthrange(y, m)[1])
    return date.replace(day=d, month=m, year=y)


next_month = monthdelta(date.today(), 1)
next_month

# %%
datetime.now().isoformat()


# %%
EST = pytz.timezone('America/New_York')
dt = parse('2014-02-03 09:17:00 EST', tzinfos={'EST': EST})
dt


# %%

dt = parse("Today is January 1, 2047 at 8:21:00AM", fuzzy=True)
print(dt)


# %%

# The size of each step in days
day_delta = datetime.timedelta(days=1)
start_date = datetime.date.today()
end_date = start_date + 7*day_delta

for i in range((end_date - start_date).days):
    print(start_date + i*day_delta)


# %%
datetime_for_string = datetime(2016, 10, 1, 0, 0)
datetime_string_format = '%b %d %Y, %H:%M:%S'
datetime.strftime(datetime_for_string, datetime_string_format)


# %%
datetime_string = 'Oct 1 2016, 00:00:00'
datetime_string_format = '%b %d %Y, %H:%M:%S'
datetime.strptime(datetime_string, datetime_string_format)


# %% [markdown]
# ## Enum

# %%


class Color(Enum):
    red = 1
    green = 2
    blue = 3


print(Color.red)
print(Color(1))
print(Color['red'])

# %%
[c for c in Color]

# %%
print({1, 2, 3, 4, 5}.intersection({3, 4, 5, 6}), {
      1, 2, 3, 4, 5} & {3, 4, 5, 6})  # Intersection
print({1, 2, 3, 4, 5}.union({3, 4, 5, 6}), {
      1, 2, 3, 4, 5} | {3, 4, 5, 6})  # Union
print({1, 2, 3, 4, 5}.difference({3, 4, 5, 6}), {
      1, 2, 3, 4, 5} - {3, 4, 5, 6})  # Difference
print({1, 2, 3, 4, 5}.symmetric_difference({3, 4, 5, 6}), {
      1, 2, 3, 4, 5} ^ {3, 4, 5, 6})  # Symmetric difference with
print({1, 2, 3, 4, 5}.issuperset({3, 4, 5, 6}), {
      1, 2, 3, 4, 5} >= {3, 4, 5, 6})  # Superset check
print({1, 2, 3, 4, 5}.issubset({3, 4, 5, 6}), {
      1, 2, 3, 4, 5} <= {3, 4, 5, 6})  # Subset check
print({1, 2, 3, 4, 5}.isdisjoint({3, 4, 5, 6}), {
      1, 2, 3, 4, 5}.isdisjoint({3, 4, 5, 6}))  # Union

# %%
print(2 in {1, 2, 3})
print(4 in {1, 2, 3})
print(4 not in {1, 2, 3})

s = {1, 2, 3}
s.add(4)  # s == {1,2,3,4}
s.discard(3)  # s == {1,2,4}
s.discard(5)  # s == {1,2,4}
s.remove(2)  # s == {1,4}
#s.remove(2) -> KeyError!


# %%

a, b = 2, 3
math.pow(a, b)


# %%
math.sqrt(2)

# %%
x = 8
math.pow(x, 1/3)

# %%
math.exp(0), math.exp(1)

# %%
math.exp(1e-6) - 1, math.expm1(1e-6)


# %%
a, b = 1, 2
math.sin(a)

# %%
math.atan(math.pi)

# %%
math.hypot(a, b)

# %% [markdown]
# * -= decrement the variable in place
# * += increment the variable in place
# * *= multiply the variable in place
# * /= divide the variable in place
# * //= floor divide the variable in place # Python 3
# * %= return the modulus of the variable in place
# * **= raise to a power in place

# %%
math.log(5)

# %%
~0

# %%
60 ^ 30
# 60 = 0b111100
# 30 = 0b011110
# 34 = 0b100010

# %%
# 60 = 0b111100
# 30 = 0b011110
60 & 30
# 28 = 0b11100

# %%
# 60 = 0b111100
# 30 = 0b011110
60 | 30
# 62 = 0b111110

# %%
#2 = 0b10
2 << 2
#8 = 0b1000

# %%
# 8 = 0b1000
8 >> 2
# Out: 2
# 2 = 0b10

# %%
a = 0b001
a &= 0b010
# a = 0b000
a = 0b001
a |= 0b010
# a = 0b011
a = 0b001
a <<= 2
# a = 0b100
a = 0b100
a >>= 2
# a = 0b001
a = 0b101
a ^= 0b011
# a = 0b110


# %%
x = True
y = True
z = x and y  # z = True
x = True
y = False
z = x and y  # z = False
x = False
y = True
z = x and y  # z = False
x = False
y = False
z = x and y  # z = False
x = 1
y = 1
z = x and y  # z = y, so z = 1, see `and` and `or` are not guaranteed to be a boolean
x = 0
y = 1
z = x and y  # z = x, so z = 0 (see above)
x = 1
y = 0
z = x and y  # z = y, so z = 0 (see above)
x = 0
y = 0
z = x and y  # z = x, so z = 0 (see above)

# %%
x = True
y = True
z = x or y  # z = True
x = True
y = False
z = x or y  # z = True
x = False
y = True
z = x or y  # z = True
x = False
y = False
z = x or y  # z = False
x = 1
y = 1
z = x or y  # z = x, so z = 1, see `and` and `or` are not guaranteed to be a boolean
x = 1
y = 0
z = x or y  # z = x, so z = 1 (see above)
x = 0
y = 1
z = x or y  # z = y, so z = 1 (see above)
x = 0
y = 0
z = x or y  # z = y, so z = 0 (see above)

# %%
x = True
y = not x  # y = False
x = False
y = not x  # y = True

# %%
x = 'hello'


def read_x():
    print(x)


read_x()

# %%
x = 5
print(x)
del x
# print(x) NameError: name 'x' is not defined

# %%
a = 'global'


class Fred:
    a = 'class'  # class scope
    b = (a for i in range(10))  # function scope
    c = [a for i in range(10)]  # function scope
    d = a  # class scope
    def e(): return a  # function scope
    def f(a=a): return a  # default argument uses class scope

    @staticmethod  # or @classmethod, or regular instance method
    def g():  # function scope
        return a


print(Fred.a)
print(next(Fred.b))
print(Fred.c[0])
print(Fred.d)
print(Fred.e())
print(Fred.f())
print(Fred.g())


# %%
foo = 0  # global foo


def f1():
    foo = 1
    # a new foo local in f1

    def f2():
        foo = 2
        # a new foo local in f2

        def f3():
            foo = 3  # a new foo local in f3
            print(foo)  # 3
            foo = 30  # modifies local foo in f3 only

            def f4():
                global foo
                print(foo)  # 0
                foo = 100  # modifies global foo


# %%
n = 5
"Greater than 2" if n > 2 else "Smaller than or equal to 2"

# %%
n = 5
"Hello" if n > 10 else "Goodbye" if n > 5 else "Good day"

# %%
number = 5
if number > 2:
    print("Number is bigger than 2.")
elif number < 2:
    print("Number is smaller than 2.")
else:
    print("Number is 2.")

# %%
1 and 2


# %%
1 and "Hello World"

# %%
print(1 > -1 < 2 > 0.5 < 100 != 24)

# %%
a = 'Python is fun!'
b = 'Python is fun!'

a == b, a is b


# %%
12 > 4, 12 != 1, 12 == 12

# %%
i = 0
while i < 7:
    print(i)
    if i == 4:
        print("Breaking from loop")
        break
    i += 1


# %%
for i in (0, 1, 2, 3, 4):
    print(i)
    if i == 2:
        break

# %%
for i in [0, 1, 2, 3, 4]:
    print(i)
for i in range(5):
    print(i)
for x in range(1, 6):
    print(x)

# %%
for index, item in enumerate(['one', 'two', 'three', 'four']):
    print(index, '::', item)


# %%
x = map(lambda e: e.upper(), ['one', 'two', 'three', 'four'])
print(list(x))

# %%
for i in range(3):
    print(i)
else:
    print('done')

i = 0
while i < 3:
    print(i)
    i += 1
else:
    print('done')


# %%
d = {"a": 1, "b": 2, "c": 3}

for key, value in d.items():
    print(key, "::", value)

# %%
collection = [('a', 'b', 'c'), ('x', 'y', 'z'), ('1', '2', '3')]
for item in collection:
    i1 = item[0]
    i2 = item[1]
    i3 = item[2]


# %%
for item in collection:
    i1, i2, i3 = item


# %%
for i1, i2, i3 in collection:
    print(i1, i2, i3)

# %%
lst = ['alpha', 'bravo', 'charlie', 'delta', 'echo']

for s in lst:
    print(s[:1])
for idx, s in enumerate(lst):
    print("%s has an index of %d" % (s, idx))


# %%
my_array = array('i', [1, 2, 3, 4])
my_array

# %%
# list like a array

# %%
lst1 = [1, 2, 3]
lst2 = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
lst3 = [[[1, 2, 3], [4, 5, 6], [7, 8, 9]], [[1, 2, 3], [
    4, 5, 6], [7, 8, 9]], [[1, 2, 3], [4, 5, 6], [7, 8, 9]]]
lst1, lst2, lst3

# %%
d = {}
d = {'key': 'value'}

# %%
d = {k: v for k, v in [('key', 'value',)]}
d

# %%
# ! [ <expression> for <element> in <iterable> if <condition> ]


# %%
squares = [x * x for x in (1, 2, 3, 4)]
squares

# %%
squares = []
for x in (1, 2, 3, 4):
    squares.append(x * x)
squares

# %%
[s.upper() for s in "Hello World"]

# %%
[w.strip(',') for w in ['these,', 'words,,', 'mostly', 'have,commas,']]

# %%
sentence = "Beautiful is better than ugly"
["".join(sorted(word, key=lambda x: x.lower())) for word in sentence.split()]


# %%
[x if x in 'aeiou' else '*' for x in 'apple']

# %%
# ! Double Iteration
[str(x) for i in range(3) for x in foo(i)]


# %%
# list comprehension
[x**2 for x in range(10)]

# %%
{x for x in range(1, 11) if x % 2 == 0}

# %%
text = "When in the Course of human events it becomes necessary for one people..."
{ch.lower() for ch in text if ch.isalpha()}

# %%
list(filter(lambda x: x % 2 == 0, range(10)))  # even numbers < 10


# %%
list(map(lambda x: 2*x, range(10)))  # multiply each number by two

# %%
lst = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h']
lst[::2], lst[::3], lst[2:4], lst[2:], lst[:4], lst[::-1]

# %%
names = ['Fred', 'Wilma', 'Barney']


def long_name(name):
    return len(name) > 5


list(filter(long_name, names))

# %%
print("You can print \n escape characters \t\t too.")


# %% [markdown]
# # File

# %%
with open('filename.txt', 'w') as object:
    object.write('merlins\n\nmerlins\n\t\t\tmerlins')
with open('filename.txt', 'r') as object:
    lines = object.readlines()
print(lines)
with open('filename.txt', 'r') as object:
    content = object.read()
    lines = content.split('\n')
print(lines)

# %%
fileobj = open('filename.txt', 'r')
pos = fileobj.tell()
print('We are at %u.' % pos)

# %%
content = fileobj.read()
end = fileobj.tell()
print('This file was %u characters long.' % end)
fileobj.close()


# %%
fileobj = open('filename.txt', 'r')
fileobj.seek(7)
pos = fileobj.tell()
print('We are at character #%u.' % pos)
next4 = fileobj.read(4)
print(next4)
pos = fileobj.tell()
print('We are at %u.' % pos)

fileobj.close()

# %%
print("Hello, ", end=" ")
print("Wolrd! ", end="\n")
print(1)

# %% [markdown]
# ## File modes
#
# * 'r' - reading mode. The default. It allows you only to read the file, not to modify it. When using this mode the file must exist.
# * 'w' - writing mode. It will create a new file if it does not exist, otherwise will erase the file and allow you to write to it.
# * 'a' - append mode. It will write data to the end of the file. It does not erase the file, and the file must exist for this mode.
# * 'rb' - reading mode in binary. This is similar to r except that the reading is forced in binary mode. This is also a default choice.
# * 'r+' - reading mode plus writing mode at the same time. This allows you to read and write into files at the same time without having to use r and w.
# * 'rb+' - reading and writing mode in binary. The same as r+ except the data is in binary
# * 'wb' - writing mode in binary. The same as w except the data is in binary.
# * 'w+' - writing and reading mode. The exact same as r+ but if the file does not exist, a new one is made. Otherwise, the file is overwritten.
# * 'wb+' - writing and reading mode in binary mode. The same as w+ but the data is in binary.
# * 'ab' - appending in binary mode. Similar to a except that the data is in binary.
# * 'a+' - appending and reading mode. Similar to w+ as it will create a new file if the file does not exist. Otherwise, the file pointer is at the end of the file if it exists.
# * 'ab+' - appending and reading mode in binary. The same as a+ except that the data is in binary.
#

# %%
os.path.join('a', 'b', 'c')

# %%
p = os.path.join(os.getcwd(), 'textFile.txt')
p

# %%
os.path.dirname(p)

# %%
os.path.basename(p)

# %%
os.path.split(os.getcwd())

# %%
os.path.splitext(os.path.basename(p))

# %%
path = '/home/john/temp'
os.path.exists(path)


# %%
dirname = '/home/merlins/Projeler/Learn_AI/python-zero-to-hero'
os.path.isdir(dirname)

# %%
filename = dirname + 'hello.py'
os.path.isfile(filename)


# %%
filename = dirname + '/hello.py'
os.path.isfile(filename)

# %%


def greet():
    print("Hello")


greet()

# %%


def many_types(x):
    if x < 0:
        return "Hello!"
    else:
        return 0


print(many_types(1))
print(many_types(-1))


# %%
def func(*args):
    for i in args:
        print(i)


func(1, 2, 3)
list_of_arg_values = [1, 2, 3]
func(*list_of_arg_values)

# %%


def func(**kwargs):
    # kwargs will be a dictionary containing the names as keys and the values as values
    for name, value in kwargs.items():
        print(name, value)


func(value1=1, value2=2, value3=3)  # Calling it with 3 arguments
func()
my_dict = {'foo': 1, 'bar': 2}
func(**my_dict)


# %%
#       |-positional-|-optional-|---keyword-only--|-optional-|
def func(arg1, arg2=10, *args, kwarg1, kwarg2=2, **kwargs):
    pass

# %%


def fn(**kwargs):
    print(kwargs)
    f1(**kwargs)


def f1(**kwargs):
    print(len(kwargs))


fn(a=1, b=2)


# %%
def strip_and_upper_case(s): return s.strip().upper()


print(strip_and_upper_case("  Hello "))

greeting = lambda x, *args, **kwargs: print(x, args, kwargs)
greeting('hello', 'world', world='world')


# %%
sorted([" foo ", " bAR", "BaZ "], key=lambda s: s.strip().upper())

# %%
sorted(map(lambda s: s.strip().upper(), [" foo ", "bAR", "BaZ "]))

# %%
sorted(map(lambda s: s.strip(), [" foo ", " bAR", "BaZ"]))

# %%
my_list = [3, -4, -2, 5, 1, 7]
sorted(my_list, key=lambda x: abs(x))

# %%
list(filter(lambda x: x > 0, my_list))

# %%
list(map(lambda x: abs(x), my_list))

# %%


def foo(msg):
    print(msg)


greet = lambda x = "hello world": foo(x)
greet()

# %%


def f(x): return 2*x


f(2)

# %%


def f(x): return 2*x


f(2)

# %%


def make(action='nothing'):
    return action


make("fun")

# %%
make(action="sleep")

# %%
make()

# %%


def f(a, b=42, c=[]):
    pass


print(f.__defaults__)

# %%


def give_me_five():
    return 5


print(give_me_five())
print(give_me_five() + 10)


def give_me_another_five():
    return 5
    print('This statement will not be printed. Ever.')


print(give_me_another_five())


def give_me_two_fives():
    return 5, 5


first, second = give_me_two_fives()
print(first, second)

# %%


def makeInc(x):
    def inc(y):
        nonlocal x
        # now assigning a value to x is allowed
        x += y
        return x
    return inc


incOne = makeInc(1)
incOne(5)

# %%


def cursing(depth):
    try:
        cursing(depth + 1)  # actually, re-cursing
    except RuntimeError as RE:
        print('I recursed {} times!'.format(depth))


cursing(0)
sys.setrecursionlimit(2000)
cursing(0)


# %%
def lambda_factorial(i): return 1 if i == 0 else i*lambda_factorial(i-1)


print(lambda_factorial(4))

# %%


def factorial(n):
    #n here should be an integer
    if n == 0:
        return 1
    else:
        return n*factorial(n-1)


factorial(4)

# %%


def factorial_2(n): return 1 if n == 0 else n*factorial_2(n-1)


factorial_2(4)

# %% [markdown]
# # Functional Programming in Python

# %%


def s(x): return x + x


s(4)

# %%
name_len = map(len, ['merlins', 'jimmy', 'Sheepman'])
list(name_len)

# %%
total = reduce(lambda a, x: a+x, [0, 1, 2, 3, 4])
total

# %%
arr = [1, 2, 3, 4, 5, 6]
[i for i in filter(lambda x:x > 4, arr)]

# %%


def super_secret_fucntion(f):
    return f


@super_secret_fucntion
def my_function():
    print("This is my secret function.")


my_function()

# %%


def print_args(func):
    def inner_func(*args, **kwargs):
        print(args)
        print(kwargs)
        # Call the original function with its arguments.
        return func(*args, **kwargs)
    return inner_func


@print_args
def multiply(num_a, num_b):
    return num_a * num_b


multiply(3, 5)

# %%


class Decorator(object):
    """Simple decorator class."""

    def __init__(self, func):
        self.func = func

    def __call__(self, *args, **kwargs):
        print('Before the function call.')
        res = self.func(*args, **kwargs)
        print('After the function call.')
        return res


@Decorator
def testfunc():
    print('Inside the function.')


testfunc()


# %%


class CountCallsDecorator(object):

    def __init__(self, func):
        self.func = func
        self.ncalls = 0

    # Number of calls of this method
    def __call__(self, *args, **kwargs):
        self.ncalls += 1
        # Increment the calls counter
        return self.func(*args, **kwargs)

    def __get__(self, instance, cls):
        return self if instance is None else MethodType(self, instance)


class Test(object):

    def __init__(self):
        pass

    @CountCallsDecorator
    def do_something(self):
        return 'something was done'


a = Test()
a.do_something()
a.do_something.ncalls
b = Test()
b.do_something()
b.do_something.ncalls

# %%


def decoratorfactory(message):
    def decorator(func):
        def wrapped_func(*args, **kwargs):
            print('The decorator wants to tell you: {}'.format(message))
            return func(*args, **kwargs)
        return wrapped_func
    return decorator


@decoratorfactory('Hello World')
def test():
    pass


test()

# %%


def decoratorfactory(*decorator_args, **decorator_kwargs):

    class Decorator(object):

        def __init__(self, func):
            self.func = func

        def __call__(self, *args, **kwargs):
            print('Inside the decorator with arguments {}'.format(decorator_args))
            return self.func(*args, **kwargs)

    return Decorator


@decoratorfactory(10)
def test():
    pass


test()


# %%
class Decorator(object):

    def __init__(self, func):
        # Copies name, module, annotations and docstring to the instance.
        self._wrapped = wraps(func)(self)

    def __call__(self, *args, **kwargs):
        return self._wrapped(*args, **kwargs)


@Decorator
def test():
    """Docstring of test."""
    pass


test.__doc__


# %%
class Person(object):
    """A simple class."""
    species = "homo sapiens"

    def __init__(self, name):
        """This is the initializer. It's a special method (see below)."""
        self.name = name

    def __str__(self):
        """This method is run when Python tries to cast the object to a string. Return
        this string when using print(), etc. """
        return self.name

    def rename(self, renamed):
        self.name = renamed
        print("Now my name is {}".format(self.name))


# %%
merlins = Person("merlins")
merlins.name

# %%
merlins.rename("jimmy")

# %%
merlins.name

# %%


class D(object):
    multiplier = 2

    @classmethod
    def f(cls, x):
        return cls.multiplier * x

    @staticmethod
    def g(name):
        print("Hello, %s" % name)


# %%
D.f

# %%
D.f(12)

# %%
D.g("hellooo")

# %%


class BaseClass(object):
    pass


class DerivedClass(BaseClass):
    pass

# %%


class Rectangle():
    def __init__(self, w, h):
        self.w = w
        self.h = h

    def area(self):
        return self.w * self.h

    def perimeter(self):
        return 2 * (self.h + self.w)

# %%


class Square(Rectangle):
    def __init__(self, s):
        super(Square, self).__init__(s, s)
        self.s = s


# %%
r = Rectangle(4, 3)

# %%
r.area()

# %%
r.perimeter()

# %%
s = Square(2)

# %%
s.area()

# %%
s.perimeter()

# %%
issubclass(Square, Rectangle)

# %%
isinstance(r, Rectangle)

# %%
isinstance(r, Square)

# %%
isinstance(s, Rectangle)

# %%

isinstance(s, Square)

# %%


class A(object):
    def __init__(self, num):
        self.num = num

    def __add__(self, other):
        return A(self.num + other.num)

# %%


def get_number(self):
    return self.num


# %%
A.get_num = get_number

# %%
foo = A(31)


# %%
foo.get_num()

# %%


class Person(object):

    def __init__(self, first_name, last_name, age):
        self.first_name = first_name
        self.last_name = last_name
        self.age = age
        self.full_name = first_name + " " + last_name

    @classmethod
    def from_full_name(cls, name, age):
        if " " not in name:
            raise ValueError
        first_name, last_name = name.split(" ", 2)
        return cls(first_name, last_name, age)

    def greet(self):
        print("Hello, my name is " + self.full_name + ".")


# %%
bob = Person("Bob", "Bobberson", 42)
alice = Person.from_full_name("Alice Henderson", 31)
bob.greet()

# %%
alice.greet()

# %%


class Foo(object):
    foo = 'attr foo of Foo'


class Bar(object):
    foo = 'attr foo of Bar'
    bar = 'attr bar of Bar'


class FooBar(Foo, Bar):
    foobar = 'attr foobar of FooBar'


# %%
fb = FooBar()
fb.foo

# %%


class MyClass(object):

    def __init__(self):
        self._my_string = ""

    @property
    def string(self):
        """A profoundly important string."""
        return self._my_string

    @string.setter
    def string(self, new_value):
        assert isinstance(new_value, str), \
            "Give me a string, not a %r!" % type(new_value)
        self._my_string = new_value

    @string.deleter
    def x(self):
        self._my_string = None


mc = MyClass()
mc.string = "String!"
print(mc.string)
# del mc.string // AttributeError: can't delete attribute

# %%


class Rectangle(object):

    def __init__(self, width, height, color='blue'):
        self.width = width
        self.height = height
        self.color = color

    def area(self):
       return self.width * self.height


default_rectangle = Rectangle(2, 3)
print(default_rectangle.color)
red_rectangle = Rectangle(2, 3, 'red')
print(red_rectangle.color)

# %%


class Singleton:
    """
    A non-thread-safe helper class to ease implementing singletons.
    This should be used as a decorator -- not a metaclass -- to the
    class that should be a singleton.
    The decorated class can define one `__init__` function that
    takes only the `self` argument. Other than that, there are
    no restrictions that apply to the decorated class.
    To get the singleton instance, use the `Instance` method. Trying
    to use `__call__` will result in a `TypeError` being raised.
    Limitations: The decorated class cannot be inherited from.
    """

    def __init__(self, decorated):
        self._decorated = decorated

    def Instance(self):
        """
        Returns the singleton instance. Upon its first call, it creates a
        new instance of the decorated class and calls its `__init__` method.
        On all subsequent calls, the already created instance is returned.
        """
        try:
            return self._instance
        except AttributeError:
            self._instance = self._decorated()
            return self._instance

    def __call__(self):
        raise TypeError('Singletons must be accessed through `Instance()`.')

    def __instancecheck__(self, inst):
        return isinstance(inst, self._decorated)


# %%
@Singleton
class Single:
    def __init__(self):
        self.name = None
        self.val = 0

    def getName(self):
        print(self.name)


x = Single.Instance()
y = Single.Instance()
x.name = 'I\'m single'
x.getName()
y.getName()


# %%
foo = 1
bar = 'bar'
baz = 3.14
print('{}, {} and {}'.format(foo, bar, baz))
print('{0}, {1}, {2}, and {1}'.format(foo, bar, baz))
# print('{0}, {1}, {2}, and {3}'.format(foo, bar, baz)) Error

# %%
print("X value is: {x_val}. Y value is: {y_val}.".format(x_val=2, y_val=3))

# %%


class AssignValue(object):
    def __init__(self, value):
        self.value = value


my_value = AssignValue(6)
print('My value is: {0.value}'.format(my_value))


# %%
my_dict = {'key': 6, 'other_key': 7}
print("My other key is: {0[other_key]}".format(my_dict))

# %%
my_list = ['zero', 'one', 'two']
print("2nd element is: {0[2]}".format(my_list))

# %%
'{:~^20}'.format('centered')

# %%
'{:~<9s}, World'.format('Hello')

# %%
'{:~>9s}, World'.format('Hello')

# %%
'{:~^9s}'.format('Hello')

# %%
'{:0=6d}'.format(-123)

# %%
foo = 'bar'

# %%
f'Foo is {foo}'

# %%
f'{foo:^7s}'

# %%
price = 478.23
f"{f'${price:0.2f}':*>20s}"

# %%
'{0:.0f}'.format(42.12345)

# %%
'{0:.5f}'.format(42.12345)

# %%
data = {'first': 'Hodor', 'last': 'Hodor!'}
'{first} {last}'.format_map(data)

# %%
'{:d}'.format(0x0a)

# %%
'{:.>10}'.format('foo')

# %%
'{:.>{}}'.format('foo', 10)

# %%
'{:{}{}{}}'.format('foo', '*', '^', 15)

# %%
print("XßΣ".casefold())
print("XßΣ".lower())
print("This is a 'string'.".upper())
print("This IS a 'string'.".lower())
print("this Is A 'String'.".capitalize())
print("this Is a 'String'".title())
print("this iS A STRiNG".swapcase())
print(str.upper("This is a 'string'"))

# %%
translation_table = str.maketrans("aeiou", "12345")
my_string = "This is a string!"
translated = my_string.translate(translation_table)
translated

# %%

print(string.ascii_letters)
print(string.ascii_lowercase)
print(string.ascii_uppercase)
print(string.digits)
print(string.hexdigits)
print(string.punctuation)
print(string.whitespace)
print(string.printable)

# %%
">>> a Python prompt".strip('> ')

# %%
[char for char in reversed('hello')]

# %%
''.join(reversed('hello'))

# %%
"This is a sentence.".split()

# %%
"This is a sentence.".split('en')

# %%
"This is a sentence.".split('e', maxsplit=2)


# %%
"Make sure to foo your sentence.".replace('foo', 'spam')

# %%


def func(params):
    for value in params:
        print('Got value {}'.format(value))
        if value == 1:
            # Returns from function as soon as value is 1
            print(">>>> Got 1")
            return
        print("Still looping")
    return "Couldn't find 1"


func([5, 3, 1, 2, 8, 9])


# %%

# %%
# PEP8 rules for Imports

# %%
x = 1.55
y = -1.55
z = 1 + 3j
w = 1-7j
pos_inf = float('inf')  # positive infinity
neg_inf = float('-inf')  # negative infinity
not_a_num = float('nan')  # NaN ("not a number")

# %%
print(round(x))
print(round(y))
print(round(x, 1))
print(round(y, 1))
print(math.floor(x))
print(math.floor(y))
print(math.ceil(x))
print(math.ceil(y))
print(math.trunc(x))
print(math.trunc(y))
print(round(1.3))
print(round(1.33, 1))
print(round(2.675, 2))
print(math.hypot(2, 4))
print(math.radians(45))
print(math.degrees(math.asin(1)))
print(math.sin(math.pi / 2))
print(math.sin(math.radians(90)))
print(math.asin(1))
print(math.asin(1) / math.pi)
print(math.cos(math.pi / 2))
print(math.acos(1))
print(math.tan(math.pi/2))
print(math.atan(math.inf))
print(math.atan(float('inf')))
print(math.atan2(1, 2))
print(math.atan2(-1, -2))
print(math.atan2(1, 0))
print(math.sinh(math.pi))
print(math.asinh(1))
print(math.cosh(math.pi))
print(math.acosh(1))
print(math.tanh(math.pi))
print(math.atanh(0.5))
print(pos_inf, neg_inf, not_a_num)

pos_inf = math.inf
neg_inf = -math.inf
not_a_num = math.nan

print(pos_inf, neg_inf, not_a_num)
print(sys.float_info.max)
print(pos_inf > sys.float_info.max)
print(neg_inf < -sys.float_info.max)
print(math.log(math.e))
print(math.log(1))
print(math.log(100))
print(math.log1p(1e-20))
print(math.log10(10))
print(math.log(100, 10))
print(z)
print(1j ** 1j)
print(z.real, z.imag)
print(z.conjugate())
print(abs(1 + 1j))
print(complex(1))

print(cmath.sqrt(-1))
print(cmath.log(1+1j))
print(cmath.polar(z))
print(cmath.rect(2, cmath.pi/2))
print(cmath.exp(z))
print(cmath.log(z))
print(cmath.log10(-100))
print(cmath.sqrt(z))
print(cmath.sin(z))
print(cmath.cos(z))
print(cmath.tan(z))
print(cmath.asin(z))
print(cmath.acos(z))
print(cmath.atan(z))
print(cmath.sin(z)**2 + cmath.cos(z)**2)
print(cmath.sinh(z))
print(cmath.cosh(z))
print(cmath.tanh(z))
print(cmath.asinh(z))
print(cmath.acosh(z))
print(cmath.atanh(z))
print(cmath.cosh(z)**2 - cmath.sin(z)**2)
print(cmath.cosh((0+1j)*z) - cmath.cos(z))
print(z+w)
print(z-w)
print(z*w)
print(z/w)
print(z**3)
print(z.real)
print(z.imag)
print(abs(z))
print(z.conjugate())

# %%
counts = collections.Counter([1, 2, 3])
counts

# %%
collections.Counter('Happy Birthday')

# %%
collections.Counter(
    'I am Sam Sam I am That Sam-I-am That Sam-I-am! I do not like that Sam-I-am'.split())

# %%
c = collections.Counter({'a': 4, 'b': 2, 'c': -2, 'd': 0})
c['a']

# %%
c['c'] = -3

# %%
c

# %%
d = deque('ghi')
for elem in d:
    print(elem.upper())

# %%
d.append('j')
d.appendleft('f')
d

# %%
d.pop()
d.popleft()
d

# %%
list(d)

# %%
list(reversed(d))

# %%
d.extend('jkl')
d

# %%
d.rotate(1)
d

# %%
d.rotate(-1)
d

# %%
deque(reversed(d))
d

# %%
d.clear()
d

# %%
# define two dictionaries with at least some keys overlapping.
dict1 = {'apple': 1, 'banana': 2}
dict2 = {'coconut': 1, 'date': 1, 'apple': 3}
# create two ChainMaps with different ordering of those dicts.
combined_dict = collections.ChainMap(dict1, dict2)
reverse_ordered_dict = collections.ChainMap(dict2, dict1)


for k, v in combined_dict.items():
    print(k, v)
print("------")
for k, v in reverse_ordered_dict.items():
    print(k, v)

# %%
adict = {'a': 1, 'b': 5, 'c': 1}
dict((i, dict(v)) for i, v in groupby(adict.items(), itemgetter(1)))

# %%
dict((i, dict(v)) for i, v in groupby(adict.items(), lambda x: x[1]))

# %%
alist_of_tuples = [(5, 2), (1, 3), (2, 2)]
sorted(alist_of_tuples, key=itemgetter(1, 0))

# %%
add(1, 1)

# %%
mul('a', 10)


# %%
mul([3], 3)

# %%
alist = ['wolf', 'sheep', 'duck']
list(filter(lambda x: x.startswith('d'), alist))

# %%
list(filter(methodcaller('startswith', 'd'), alist))

# %%
d = {
    'foo': 'bar',
    'alice': 1,
    'wonderland': [1, 2, 3]
}

with open('text.txt', 'w') as f:
    json.dump(d, f)

# %%
with open('text.txt', 'r') as f:
    d = json.load(f)

# %%
data = {"cats": [{"name": "Tubbs", "color": "white"},
                 {"name": "Pepper", "color": "black"}]}

# %%
print(json.dumps(data))

# %%
print(json.dumps(data, indent=2))

# %%
print(json.dumps(data, sort_keys=True))

# %%
print(json.dumps(data, separators=(',', ':')))

# %%
json_file_path = './data.json'
data = {u"foo": u"bar", u"baz": []}

with open(json_file_path, 'w') as json_file:
    json.dump(data, json_file)

with open(json_file_path) as json_file:
    json_file_content = json_file.read()

with open(json_file_path) as json_file:
    json.load(json_file)

# %%
d = {
    'foo': 'bar',
    'alice': 1,
    'wonderland': [1, 2, 3]
}
json.dumps(d)

# %%
s = '{"wonderland": [1, 2, 3], "foo": "bar", "alice": 1}'
json.loads(s)

# %%

connection = sqlite3.connect('example.db')

commender = connection.cursor()

commender.execute(
    '''create table stocks(date text, trans text,symbol text,qty real,price real)''')

commender.execute(
    "insert into stocks values('2006-01-05','BUY','RHAT',100,35.14)")

connection.commit()
connection.close()

# %%
