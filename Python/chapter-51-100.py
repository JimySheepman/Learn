# %% [markdown]
# # chapter 51-100

# %%
from requests import post
from bs4 import BeautifulSoup
import urllib
import urllib.request
import re
import pprint
from operator import mul
import operator
import matplotlib.pyplot as plt
import numpy as np
import webbrowser
import base64
from functools import reduce
import functools
from functools import partial
import random
from random import SystemRandom
import itertools
from itertools import *
import locale
import os

# %%
os.mkdir('newdir')


# %%
print(os.getcwd())

# %%
os.name

# %%
os.rmdir('newdir')


# %%

# %%
locale.setlocale(locale.LC_ALL, '')

# %%
locale.currency(6512.1)

# %%
locale.currency(762559748.49, grouping=True)

# %%
a = [1, 2, 3, 4, 5]
b = list(itertools.combinations(a, 2))


# %%
a = [i for i in range(5)]
b = ['a', 'b', 'c', 'd', 'e', 'f', 'g']
for i in zip_longest(a, b):
    x, y = i
    print(x, y)


# %%
def gen():
    n = 0
    while n < 20:
        n += 1
        yield n


for part in itertools.islice(gen(), 3):
    print(part)


# %%
lst = [("a", 5, 6), ("b", 2, 4), ("a", 2, 5), ("c", 2, 6)]


def testGroupBy(lst):
    groups = itertools.groupby(lst, key=lambda x: x[1])
    for key, group in groups:
        print(key, list(group))


testGroupBy(lst)

# %%
secure_rand_gen = SystemRandom()
print([secure_rand_gen.randrange(10) for i in range(10)])

# %%
print(secure_rand_gen.randint(0, 20))

# %%
laughs = ["Hi", "Ho", "He"]
random.shuffle(laughs)
laughs

# %%
random.randint(1, 8)

# %%
random.randrange(20, 50)

# %%
random.seed(5)  # Reset the random module to the same fixed state.
print(random.randrange(0, 10))


# %%
unhex = partial(int, base=16)
unhex.__doc__ = 'Convert base16 string to int'
unhex('ca11ab1e')

# %%


def f(a, b, c, x):
    return 1000*a + 100*b + 10*c + x


# %%
g = partial(f, 1, 1, 1)

# %%
print(g(2))

# %%
sorted(["A", "S", "F", "D"], key=functools.cmp_to_key(locale.strcoll))

# %%


def factorial(n):
    return reduce(lambda a, b: (a*b), range(1, n+1))


factorial(12)

# %%


# %%
s = "Hello World!"
b = s.encode("UTF-8")
s, b

# %%
e = base64.b64encode(b)
print(e)

# %%
s = "Hello World!"
b = s.encode("UTF-8")
e = base64.b64encode(b)
s1 = e.decode("UTF-8")
print(s1)

# %%
s = "Hello World!"
b = s.encode("UTF-8")
e = base64.b64encode(b)
s1 = e.decode("UTF-8")
print("Base64 Encoded:", s1)
b1 = s1.encode("UTF-8")
d = base64.b64decode(b1)
s2 = d.decode("UTF-8")
print(s2)

# %%
s = "Hello World!"
b = s.encode("UTF-8")
e = base64.b32encode(b)
s1 = e.decode("UTF-8")
print("Base32 Encoded:", s1)
b1 = s1.encode("UTF-8")
d = base64.b32decode(b1)
s2 = d.decode("UTF-8")
print(s2)

# %%
s = "Hello World!"
b = s.encode("UTF-8")
e = base64.b16encode(b)
s1 = e.decode("UTF-8")
print("Base16 Encoded:", s1)
b1 = s1.encode("UTF-8")
d = base64.b16decode(b1)
s2 = d.decode("UTF-8")
print(s2)

# %%
s = "Hello World!"
b = s.encode("UTF-8")
e = base64.a85encode(b)
s1 = e.decode("UTF-8")
print("ASCII85 Encoded:", s1)
b1 = s1.encode("UTF-8")
d = base64.a85decode(b1)
s2 = d.decode("UTF-8")
print(s2)

# %%
s = "Hello World!"
b = s.encode("UTF-8")
e = base64.b85encode(b)
s1 = e.decode("UTF-8")
print("Base85 Encoded:", s1)
b1 = s1.encode("UTF-8")
d = base64.b85decode(b1)
s2 = d.decode("UTF-8")
print(s2)

# %%
webbrowser.open("http://stackoverflow.com")
webbrowser.open_new("http://stackoverflow.com")
webbrowser.open_new_tab("http://stackoverflow.com")


# %%

# %%
x = np.linspace(0, 2.0*np.pi, 101)
y = np.sin(x)
z = np.sinh(x)
fig, ax1 = plt.subplots()
ax2 = ax1.twinx()
curve1, = ax1.plot(x, y, label="sin", color='r')
curve2, = ax2.plot(x, z, label="sinh", color='b')
curves = [curve1, curve2]
ax1.legend(curves, [curve.get_label() for curve in curves])
plt.title("Plot of sine and hyperbolic sine")
plt.show()

# %%
y = np.linspace(0, 2.0*np.pi, 101)
x1 = np.sin(y)
x2 = np.sinh(y)
ynumbers = np.linspace(0, 7, 15)
xnumbers1 = np.linspace(-1, 1, 11)
xnumbers2 = np.linspace(0, 300, 7)
fig, ax1 = plt.subplots()
ax2 = ax1.twiny()
curve1, = ax1.plot(x1, y, label="sin", color='r')
curve2, = ax2.plot(x2, y, label="sinh", color='b')
curves = [curve1, curve2]
ax2.legend(curves, [curve.get_label() for curve in curves])
ax1.set_xlabel("Magnitude", color=curve1.get_color())
ax2.set_xlabel("Magnitude", color=curve2.get_color())
ax1.set_ylabel("Angle/Value", color=curve1.get_color())
ax1.tick_params(axis='y', colors=curve1.get_color())
ax1.tick_params(axis='x', colors=curve1.get_color())
ax2.tick_params(axis='x', colors=curve2.get_color())
ax1.set_xticks(xnumbers1)
ax2.set_xticks(xnumbers2)
ax1.set_yticks(ynumbers)
ax1.grid(color=curve2.get_color())
ax2.grid(color=curve2.get_color())
ax1.xaxis.grid(False)
plt.title("Plot of sine and hyperbolic sine")
plt.show()

# %%
x = np.linspace(0, 2.0*np.pi, 101)
y = np.sin(x)
plt.plot(x, y)
plt.show()

# %%
x = np.linspace(0, 2.0*np.pi, 101)
y = np.sin(x)
xnumbers = np.linspace(0, 7, 15)
ynumbers = np.linspace(-1, 1, 11)
plt.plot(x, y, color='r', label='sin')
plt.xlabel("Angle in Radians")
plt.ylabel("Magnitude")
plt.title("Plot of some trigonometric functions")
plt.xticks(xnumbers)
plt.yticks(ynumbers)
plt.legend()
plt.grid()
plt.axis([0, 6.5, -1.1, 1.1])
plt.show()

# %%
x = np.linspace(0, 2.0*np.pi, 101)
y = np.sin(x)
z = np.cos(x)
xnumbers = np.linspace(0, 7, 15)
ynumbers = np.linspace(-1, 1, 11)
plt.plot(x, y, 'r', x, z, 'g')
plt.xlabel("Angle in Radians")
plt.ylabel("Magnitude")
plt.title("Plot of some trigonometric functions")
plt.xticks(xnumbers)
plt.yticks(ynumbers)
plt.legend(['sine', 'cosine'])
plt.grid()
plt.axis([0, 6.5, -1.1, 1.1])
plt.show()

# %%
x = np.linspace(0, 2.0*np.pi, 101)
y = np.sin(x)
z = np.cos(x)
xnumbers = np.linspace(0, 7, 15)
ynumbers = np.linspace(-1, 1, 11)
plt.plot(x, y, color='r', label='sin')
plt.plot(x, z, color='g', label='cos')
plt.xlabel("Angle in Radians")
plt.ylabel("Magnitude")
plt.title("Plot of some trigonometric functions")
plt.xticks(xnumbers)
plt.yticks(ynumbers)
plt.legend()
plt.grid()
plt.axis([0, 6.5, -1.1, 1.1])
plt.show()

# %%
expression = (x**2 for x in range(10))
list(expression)

# %%


def function():
    for x in range(10):
        yield x**2


list(function())

# %%

# %%


def add(s1, s2):
    return s1 + s2


asequence = [1, 2, 3]

reduce(add, asequence)

# %%
reduce(operator.add, asequence)

# %%


def multiply(s1, s2):
    print('{arg1} * {arg2} = {res}'.format(arg1=s1, arg2=s2, res=s1*s2))
    return s1 * s2


asequence = [1, 2, 3]

# %%
cumprod = reduce(multiply, asequence, 5)

# %%
print(cumprod)

# %%
cumprod = reduce(multiply, asequence)
print(cumprod)

# %%
reduce(operator.mul, [10, 5, -3])

# %%
reduce(operator.and_, [False, True, True, True])

# %%
reduce(operator.or_, [True, False, False, False])

# %%
names = ['Fred', 'Wilma', 'Barney']
list(map(len, names))

# %%
list(map(abs, (1, -1, 2, -2, 3, -3)))

# %%
list(map(lambda x: x*2, [1, 2, 3, 4, 5]))

# %%


def to_percent(num):
    return num * 100


list(map(to_percent, [0.95, 0.75, 1.01, 0.1]))

# %%
rate = 8.7  # fictitious exchange rate, 1 dollar = 8.7 tl
dollars = {'under_my_bed': 1000,
           'jeans': 45,
           'bank': 5000}
sum(map(partial(mul, rate), dollars.values()))


# %%
def average(*args):
    return float(sum(args)) / len(args)


measurement1 = [100, 111, 99, 97]
measurement2 = [102, 117, 91, 102]
measurement3 = [104, 102, 95, 101]
list(map(average, measurement1, measurement2, measurement3))


# %%
carnivores = ['lion', 'tiger', 'leopard', 'arctic fox']
herbivores = ['african buffalo', 'moose', 'okapi', 'parakeet']
omnivores = ['chicken', 'dove', 'mouse', 'pig']
insects = ['fly', 'ant', 'beetle', 'cankerworm']


def animals(w, x, y, z):
    return '{0}, {1}, {2}, and {3} ARE ALL ANIMALS'.format(w.title(), x, y, z)


pprint.pprint(list(map(animals, insects, carnivores, herbivores, omnivores)))

# %%
alist = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
5 in alist

# %%
atuple = ('0', '1', '2', '3', '4')
4 in atuple

# %%
astring = 'i am a string'
'a' in astring

# %%
aset = {(10, 10), (20, 20), (30, 30)}
(10, 10) in aset

# %%
adict = {0: 'a', 1: 'b', 2: 'c', 3: 'd'}
1 in adict

# %%


class ListList:
    def __init__(self, value):
        self.value = value
        self.setofvalues = set(
            item for sublist in self.value for item in sublist)

    def __iter__(self):
        print('Using __iter__.')
        return (item for sublist in self.value for item in sublist)

    def __contains__(self, value):
        print('Using __contains__.')
        return value in self.setofvalues


# %%
a = ListList([[1, 1, 1], [0, 1, 1], [1, 5, 1]])
10 in a

# %%
del ListList.__contains__
5 in a

# %%
astring = 'Hello on StackOverflow'
astring.index('o')

# %%
astring.rindex('o')

# %%
alist = [10, 16, 26, 5, 2, 19, 105, 26]
alist.index(16)

# %%
pattern = r"123"
string = "123zzb"

re.match(pattern, string)


# %%
match = re.match(pattern, string)
match.group()

# %%
pattern = r"(your base)"
sentence = "All your base are belong to us."
match = re.search(pattern, sentence)

match.group(1)
match = re.search(r"(belong.*)", sentence)
match.group(1)

# %%
match = re.search(r"^123", "123zzb")
match.group(0)

# %%
precompiled_pattern = re.compile(r"(\d+)")
matches = precompiled_pattern.search("The answer is 41!")
matches.group(1)

# %%
matches = precompiled_pattern.search("Or was it 42?")
matches.group(1)

# %%
precompiled_pattern = re.compile(r"(.*\d+)")
matches = precompiled_pattern.match("The answer is 41!")
print(matches.group(1))
matches = precompiled_pattern.match("Or was it 42?")
print(matches.group(1))

# %%
re.sub(r"t[0-9][0-9]", "foo", "my name t13 is t44 what t99 ever t44")

# %%
re.sub(r"t([0-9])([0-9])", r"t\2\1", "t13 t19 t81 t25")

# %%
re.sub(r"t([0-9])([0-9])", r"t\g<2>\g<1>", "t13 t19 t81 t25")

# %%
items = ["zero", "one", "two"]
re.sub(r"a\[([0-3])\]", lambda match: items[int(match.group(1))],
       "Items: a[0], a[1], something,a[2]")


# %%
results = re.finditer(r"([0-9]{2,3})", "some 1 text 12 is 945 here 4445588899")
print(results)

for result in results:
    print(result.group(0))


# %%
def is_allowed(string):
    characherRegex = re.compile(r'[^a-zA-Z0-9.]')
    string = characherRegex.search(string)
    return not bool(string)


print(is_allowed("abyzABYZ0099"))
print(is_allowed("#*@#$%^"))


# %%
data = re.split(r'\s+', 'James 94 Samantha 417 Scarlett 74')
print(data)

# %%
d1 = {1: []}
d2 = d1.copy()
d1 is d2

# %%
d1[1] is d2[1]

# %% [markdown]
# # Chapter 80: Creating Python packages read

# %%


def factorial(n):
    if n == 0:
        return 1
    elif n == 1:
        return 1
    else:
        return n * factorial(n - 1)


factorial(5)


# %%
def fib(n):
    if n == 0 or n == 1:
        return n
    else:
        return fib(n - 2) + fib(n - 1)


fib(5)

# %%


def two_sum(a: int, b: int):
    return a + b


def two_sum_2(a: int, b: int) -> int:
    return a + b


print(two_sum(2, 4))
print(two_sum_2(0, 0))

# %%
try:
    x = 5 / 0
except ZeroDivisionError as e:
    print("Got a divide by zero! -> ", e)
    x = 0
finally:
    print('end')

print(x)

# %%
try:
    d = {}
    a = d[1]
    b = d.non_existing_field
except KeyError as e:
    print("A KeyError has occurred. Exception message:", e)
except AttributeError as e:
    print("An AttributeError has occurred. Exception message:", e)

# %%
try:
    data = {1: 'one', 2: 'two'}
    print(data[1])
except KeyError as e:
    print('key not found')
else:
    raise ValueError()


# %%
def even_the_odds(odds):
    if odds % 2 != 1:
        raise ValueError("Did not get an odd number")
    return odds + 1


even_the_odds(3)

# %%
even_the_odds(4)

# %%

print(urllib.request.urlopen("http://stackoverflow.com/documentation/"))


# %%
response = urllib.request.urlopen("http://stackoverflow.com/documentation/")
print(response.code)

print(response.read()[:550])


# %%
query_parms = {'username': 'stackoverflow', 'password': 'me.me'}
encoded_parms = urllib.parse.urlencode(query_parms).encode('utf-8')
response = urllib.request.urlopen(
    "https://stackoverflow.com/users/login", encoded_parms)
response.code

# %%
response.read()

# %%
response = urllib.request.urlopen("http://stackoverflow.com/")
data = response.read()
encoding = response.info().get_content_charset()
html = data.decode(encoding)


# %%


# %%
data = """
<ul>
<li class="item">item1</li>
<li class="item">item2</li>
<li class="item">item3</li>
</ul>
"""
soup = BeautifulSoup(data, "html.parser")
for item in soup.select("li.item"):
    print(item.get_text())


# %%
data = """
<div>
<label>Name:</label>
John Smith
</div>
"""
soup = BeautifulSoup(data, "html.parser")
label = soup.find("label", text="Name:")
print(label.next_sibling.strip())

# %%
foo = post('http://httpbin.org/post', data={'key': 'value'})

# %%
print(foo.headers)

# %%
foo = post('http://httpbin.org/post', data={'key': 'value'}, verify=False)

# %%
print(foo.headers)

# %%
foo = post('http://httpbin.org/post',
           data={'key': 'value'}, allow_redirects=False)
print(foo.url)
print(foo.history)

# %%
payload = {'key1': 'value1',
           'key2': 'value2'
           }
foo = post('http://httpbin.org/post', data=payload)
foo.cookies

# %%


class A:
    def __init__(self, a):
        self.a = a

    def __add__(self, other):
        return self.a + other

    def __radd__(self, other):
        print("radd")
        return other + self.a


A(1) + 2

# %%
2 + A(1)


# %%
class Parent(object):
    def introduce(self):
        print("Hello!")

    def print_name(self):
        print("Parent")


class Child(Parent):
    def print_name(self):
        print("Child")


# %%
p = Parent()
c = Child()
p.introduce()
p.print_name()
c.introduce()
c.print_name()
