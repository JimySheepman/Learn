class Node:
 
    def __init__(self, key):
        self.key = key
        self.left = None
        self.right = None
 

def inorder(root):
    if root is not None:
        inorder(root.left)
        print(root.key, end=" ")
        inorder(root.right)
 

def insert(node, key):

    if node is None:
        return Node(key)
 
    if key < node.key:
        node.left = insert(node.left, key)
    else:
        node.right = insert(node.right, key)
 
    return node
 
def deleteNode(root, key):
 
    # Base Case
    if root is None:
        return root
 
    if key < root.key:
        root.left = deleteNode(root.left, key)
        return root
 
    elif(key > root.key):
        root.right = deleteNode(root.right, key)
        return root

    if root.left is None and root.right is None:
          return None
  
    if root.left is None:
        temp = root.right
        root = None
        return temp
 
    elif root.right is None:
        temp = root.left
        root = None
        return temp
 
    succParent = root
    succ = root.right
 
    while succ.left != None:
        succParent = succ
        succ = succ.left

    if succParent != root:
        succParent.left = succ.right
    else:
        succParent.right = succ.right
 
    root.key = succ.key
 
    return root

if __name__ == '__main__':
    root = None
    root = insert(root, 50)
    root = insert(root, 30)
    root = insert(root, 20)
    root = insert(root, 40)
    root = insert(root, 70)
    root = insert(root, 60)
    root = insert(root, 80)
    
    print("Inorder traversal of the given tree")
    inorder(root)
    
    print("\nDelete 20")
    root = deleteNode(root, 20)
    print("Inorder traversal of the modified tree")
    inorder(root)
    
    print("\nDelete 30")
    root = deleteNode(root, 30)
    print("Inorder traversal of the modified tree")
    inorder(root)
    
    print("\nDelete 50")
    root = deleteNode(root, 50)
    print("Inorder traversal of the modified tree")
    inorder(root)
 
