import string

def encrypt(plaintext, key):
    result = ""
    i = 0
    plaintext = makePlaintext(plaintext)
    key = makeKey(key)
    
    while (i < len(plaintext)):
        a = plaintext[i]
        b = plaintext[i+1]
        if (isSameX(a, b, key)):
            result += getRight(a, key)
            result += getRight(b, key)
        elif (isSameY(a, b, key)):
            result += getDown(a, key)
            result += getDown(b, key)
        else:
            result += getPotongan(a, b, key)
        i += 2
    
    return result

def decrypt(ciphertext, key):
    result = ""
    i = 0
    key = makeKey(key)
    
    while (i < len(ciphertext)):
        a = ciphertext[i]
        b = ciphertext[i+1]
        if (isSameX(a, b, key)):
            result += getLeft(a, key)
            result += getLeft(b, key)
        elif (isSameY(a, b, key)):
            result += getUp(a, key)
            result += getUp(b, key)
        else:
            result += getPotongan(a, b, key)
        i += 2
    
    return result

def makePlaintext(text):
    text = text.replace('j', 'i')
    temp = text[0]
    result = text[0]
    for i in range(1, len(text)):
        if (text[i] == temp):
            if (text[i] == 'x'):
                result += 'z'
            else:
                result += 'x'
        result += text[i]
        temp = text[i]
    if (len(result) % 2 != 0):
        result += 'x'
    return result

def makeKey(key):
    result = ""

    key = key.replace(" ", "")
    key = ''.join(i for i in key if not i.isdigit())
    key = key.translate(str.maketrans('', '', string.punctuation))
    key = key.lower()
    
    for char in key:
        if ((char != "j") and (result.find(char) == -1)):
            result += char
    for i in range(26):
        char = chr(i + 97)
        if ((char != "j") and (result.find(char) == -1)):
            result += char
    return result

def makeKeyMatrix(key):
    text_matrix = []
    for i in range(5):
        text_matrix.append([])
        for j in range(5):
            text_matrix[i].append(key[5 * i + j])
        print(text_matrix[i])

def isSameX(a, b, key):
    return ((key.find(a) // 5) == (key.find(b) // 5))
    
def isSameY(a, b, key):
    return ((key.find(a) % 5) == (key.find(b) % 5))

def getRight(c, key):
    if (key.find(c) % 5 == 4):
        return key[key.find(c) - 4]
    else:
        return key[key.find(c) + 1]

def getDown(c, key):
    if (key.find(c) // 5 == 4):
        return key[key.find(c) % 5]
    else:
        return key[key.find(c) + 5]

def getLeft(c, key):
    if (key.find(c) % 5 == 0):
        return key[key.find(c) + 4]
    else:
        return key[key.find(c) - 1]

def getUp(c, key):
    if (key.find(c) // 5 == 0):
        return key[key.find(c) + 20]
    else:
        return key[key.find(c) - 5]

def getPotongan(a, b, key):
    x_a = key.find(a) % 5
    y_a = key.find(a) // 5
    x_b = key.find(b) % 5
    y_b = key.find(b) // 5
    return (key[y_a * 5 + x_b] + key[y_b * 5 + x_a])