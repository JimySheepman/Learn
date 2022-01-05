# %%
import seaborn as sns
import numpy as np
import json
from struct import pack
import pickle
import sys
import matplotlib.pyplot as plt
import csv
import pandas as pd


class Card:
    def __init__(self, suit, pips):
        self.suit = suit
        self.pips = pips


# %%
ace_of_spades = Card('Spades', 1)
four_of_clubs = Card('Clubs', 4)
six_of_hearts = Card('Hearts', 6)

# %%
my_hand = [ace_of_spades, four_of_clubs, six_of_hearts]
print(my_hand)

# %%
string_of_card = str(ace_of_spades)
print(string_of_card)

# %%


class Card:
    def __init__(self, suit, pips):
        self.suit = suit
        self.pips = pips

    def __str__(self):
        special_names = {1: 'Ace', 11: 'Jack', 12: 'Queen', 13: 'King'}
        card_name = special_names.get(self.pips, str(self.pips))
        return "%s of %s" % (card_name, self.suit)


# %%
ace_of_spades = Card('Spades', 1)
print(ace_of_spades)

# %%


class Card:
    special_names = {1: 'Ace', 11: 'Jack', 12: 'Queen', 13: 'King'}

    def __init__(self, suit, pips):
        self.suit = suit
        self.pips = pips

    def __repr__(self):
        card_name = Card.special_names.get(self.pips, str(self.pips))
        return "%s of %s" % (card_name, self.suit)


# %%
print(six_of_hearts)
print(str(six_of_hearts))

# %%
print([six_of_hearts])

# %%
print(repr(six_of_hearts))

# %%
d = {'a': (1, 101), 'b': (2, 202), 'c': (3, 303)}
df = pd.DataFrame.from_dict(d, orient="index")
df.to_csv("data.csv")

# %%
df = pd.read_csv("data.csv")
d = df.to_dict()
d

# %%

with open('output.tsv', 'wt') as out_file:
    tsv_writer = csv.writer(out_file, delimiter='\t')
    tsv_writer.writerow(['name', 'field'])
    tsv_writer.writerow(['Dijkstra', 'Computer Science'])
    tsv_writer.writerow(['Shelah', 'Math'])
    tsv_writer.writerow(['Aumann', 'Economic Sciences'])

# %%
code = compile('a * b + c', '<string>', 'eval')
code

# %%
a, b, c = 1, 2, 3
eval(code)

# %%
data = np.random.randn(1000)
sns.distplot(data, kde=True, rug=True)

# %%
sns.set_style('dark')
sns.distplot(data, kde=True, rug=True)

# %%

sns.distplot(data, kde=True, rug=True)
plt.xlabel('This is my x-axis')
plt.ylabel('This is my y-axis')

# %%
x = [0, 1, 2, 3, 4, 5, 6]
y = [i**2 for i in x]
plt.scatter(x, y, c='blue', marker='x', s=100)
plt.plot(x, y, color='red', linewidth=2)
plt.xlabel('x data')
plt.ylabel('y data')
plt.title('An example plot')
plt.show()

# %%


def print_kwargs(**kwargs):
    print(kwargs)


print_kwargs(a="two", b=3)

# %%


def example(a, **kw):
    print(kw)


example(a=2, b=3, c=4)

# %%


def print_kwargs(**kwargs):
    for key in kwargs:
        print("key = {0}, value = {1}".format(key, kwargs[key]))


print_kwargs(a="two", b=1)

# %%


def print_args(farg, *args):
    print("formal arg: %s" % farg)
    for arg in args:
        print("another positional arg: %s" % arg)


print_args(1, "two", 3)

# %%


def foobar(foo=None, bar=None):
    return "{}{}".format(foo, bar)


values = {"foo": "foo", "bar": "bar"}
foobar(**values)

# %%

sys.getrefcount(1)

# %%
a = 999999999

sys.getrefcount(999999999)


# %%

data = {
    'a': [1, 2.0, 3, 4+6j],
    'b': ("character string", b"byte string"),
    'c': {None, True, False}
}
with open('data.pickle', 'wb') as f:
    pickle.dump(data, f, pickle.HIGHEST_PROTOCOL)


# %%
with open('data.pickle', 'rb') as f:
    data_1 = pickle.load(f)

# %%
data_1

# %%

print(pack('I3c', 123, b'a', b'b', b'c'))

# %%
x = True
y = False
x, y = y, x
x, y

# %%
families = (['John'], ['Mark', 'David', {'name': 'Avraham'}])
# Dumping it into string
json_families = json.dumps(families)
# [["John"], ["Mark", "David", {"name": "Avraham"}]]
# Dumping it to file
with open('families.json', 'w') as json_file:
    json.dump(families, json_file)
# Loading it from string
json_families = json.loads(json_families)
# Loading it from file
with open('families.json', 'r') as json_file:
    json_families = json.load(json_file)

# %%
json_families
