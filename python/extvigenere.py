import random
import string

# Repeat the keyword until equal length with plaintext
def repeatKey(plaintext, key):
    key = list(key)
    if len(plaintext) == len(key):
        return(key)
    else:
        for i in range(len(plaintext) - len(key)):
            key.append(key[i % len(key)])
    return(''.join(str (k) for k in key))

    
# MAIN FUNCTIONS

# Extended Vigenere Cipher
def encrypt(plaintext, key):
    key = repeatKey(plaintext, key)
    ciphertext = []
    for i in range(len(plaintext)): 
        x = ord(plaintext[i]) + ord(key[i]) % 256
        ciphertext.append(chr(x)) 
    return("" . join(ciphertext))

def decrypt(ciphertext, key):
    key = repeatKey(ciphertext, key)
    plaintext = [] 
    for i in range(len(ciphertext)):
        x = (ord(ciphertext[i]) - ord(key[i]) + 256) % 256 
        plaintext.append(chr(x)) 
    return("" . join(plaintext))
