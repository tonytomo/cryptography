import numpy as np
from sympy import Matrix

def encrypt(plaintext, matrix):
    result = ""
    m = len(matrix)
    matrix = toIntMatrix(matrix)

    plaintext_matrix = textToMatrix(plaintext, m)

    for chunk in plaintext_matrix:
        temp_matrix = np.matmul(matrix, chunk)
        for character in temp_matrix:
            result += chr(character % 26 + 97)
    
    return result

def decrypt(ciphertext, matrix):
    result = ""
    m = len(matrix)
    matrix = toIntMatrix(matrix)

    plaintext_matrix = textToMatrix(ciphertext, m)

    inv_matrix = Matrix(matrix)
    inv_matrix = inv_matrix.inv_mod(26)

    for chunk in plaintext_matrix:
        temp_matrix = np.matmul(inv_matrix, chunk)
        for character in temp_matrix:
            result += chr(character % 26 + 97)

    return result

def toIntMatrix(matrix):
    for i in range(len(matrix)):
        for j in range(len(matrix[i])):
            matrix[i][j] = int(matrix[i][j])
    
    return matrix

def textToMatrix(text, m):
    text_matrix = []
    size = m

    for adds in range(len(text) % m + 1):
        text += 'z'

    for i in range(len(text) // m):
        text_matrix.append([])
        for j in range(m):
            text_matrix[i].append((ord(text[m * i + j]) - 97) % 26)
    return(text_matrix)
