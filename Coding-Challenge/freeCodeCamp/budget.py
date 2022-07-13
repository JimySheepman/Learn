class Category:
    def __init__(self, name):
        self.name = name
        self.ledger = []
        self.balance = 0.0


    def __str__(self):
        title = f"{self.name.center(30, '*')}\n"
        body = ""
        for transaction in self.ledger:
            left_content = "{:<23}".format(transaction['description'])
            right_content = "{:>7.2f}".format(transaction['amount'])
            body += "{0}{1}\n".format(left_content[:23], right_content[:7])
        total = "Total: {:.2f}".format(self.balance)
        output = title + body + total
        return output

    def deposit(self, amount, description=""):
        self.ledger.append({
            "amount": amount,
            "description": description,
        })
        self.balance += amount



    def withdraw(self, amount, description=""):
        if self.balance - amount >=0:
            self.ledger.append({
                "amount": amount * -1,
                "description": description,
            })
            self.balance -= amount
            return True
        return False


    def get_balance(self):
        return self.balance

    def transfer(self, amount, category):
        if self.withdraw(amount, description=f"Transfer to {category.name}"):
            category.deposit(amount, description=f"Transfer from {self.name}")
            return True
        else:
            return False

    def check_funds(self, amount):
        if self.balance >= amount:
            return True
        else:
            return False


def create_spend_chart(categories):
    spend_amounts = []
    for category in categories:
        amount_spend = 0
        for transaction in category.ledger:
            if transaction['amount'] < 0:
                amount_spend += abs(transaction['amount'])
        spend_amounts.append(round(amount_spend, 2))

    total = round(sum(spend_amounts), 2)
    spend_percentage = list(map(lambda amount: int((((amount / total) * 10) // 1) * 10), spend_amounts))

    header = "Percentage spent by category\n"
    chart = ""
    for value in reversed(range(0, 101, 10)):
        chart += str(value).rjust(3) + '|'
        for percent in spend_percentage:
            if percent >= value:
                chart += " o "
            else:
                chart += "   "
        chart += " \n"

    footer = "    " + "-" * ((3 * len(categories)) + 1) + "\n"
    labels = list(map(lambda category: category.name, categories))
    max_length = max(map(lambda  label: len(label), labels))
    labels = list(map(lambda label: label.ljust(max_length), labels))
    for x in zip(*labels):
        footer += "    " + "".join(map(lambda s: s.center(3), x)) + " \n"

    return (header + chart + footer).rstrip("\n")
