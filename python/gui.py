import tkinter as tk
from tkinter import Label, StringVar, filedialog
import hill, playfair, extvigenere
import string

fields_text = 'Text', 'Key'
fields_menu = 'Pilih Input', 'Pilih Aksi', 'Pilih Cipher'

Pilih_Input = [
    "Text",
    "File Text",
    "File Random"
] 

Pilih_Aksi = [
    "Encrypt",
    "Decrypt"
]

Pilih_Cipher = [
    "Extended Vigenere Cipher",
    "Playfair Cipher",
    "Hill Cipher"
]

result = ""
spaced_result = ""
filename = ""

def fetch(entries):
    # for entry in entries:
    #     field = entry[0]
    #     text  = entry[1].get()
    #     print('%s: "%s"' % (field, text))
    # print(getKey(entries), getText(entries), entries[0][1].get(), entries[3][1].get())

    res = ""

    if (entries[2][1].get() == "Playfair Cipher"):
        if (entries[1][1].get() == "Encrypt"):
            res = playfair.encrypt(getText(entries), getKey(entries))
        elif (entries[1][1].get() == "Decrypt"):
            res = playfair.decrypt(getText(entries), getKey(entries))

    elif (entries[2][1].get() == "Extended Vigenere Cipher"):
        if (entries[1][1].get() == "Encrypt"):
            res = extvigenere.encrypt(getText(entries), getKey(entries))
        elif (entries[1][1].get() == "Decrypt"):
            res = extvigenere.decrypt(getText(entries), getKey(entries))

    elif (entries[2][1].get() == "Hill Cipher"):
        if (entries[1][1].get() == "Encrypt"):
            res = hill.encrypt(getText(entries), getKey(entries))
        elif (entries[1][1].get() == "Decrypt"):
            res = hill.decrypt(getText(entries), getKey(entries))

    printResult(res)

def printResult(res):
    if ((entries[0][1].get() == "Text") or (entries[0][1].get() == "File Text")):
        result.set(res)
        if (entries[1][1].get() == "Encrypt"):
            spaced_result.set(getSpacedResult(res))
            printResultFile(res) #print ke file
        elif (entries[1][1].get() == "Decrypt"):
            spaced_result.set("")
    else:
        if (entries[1][1].get() == "Encrypt"):
            printEncryptByteFile(res)
            result.set("Result saved to Encrypt.txt")
            spaced_result.set("")
        elif (entries[1][1].get() == "Decrypt"):
            printDecryptByteFile(res)
            result.set("Result saved to Decrypt, add formatting accordingly")
            spaced_result.set("")

def printResultFile(res):
    text_file = open("Encrypt.txt", "w", encoding="utf-8")
    text_file.write(res)
    text_file.close()

def printEncryptByteFile(res):
    text_file = open("Encrypt.txt", "wb")
    text_file.write(res.encode('utf-8'))
    text_file.close()

def printDecryptByteFile(res):
    text_file = open("Decrypt", "wb")
    text_file.write(res.encode('latin-1'))
    text_file.close()

def getText(entries):
    if ((entries[0][1].get() == "Text") or (entries[0][1].get() == "File Text")):
        # Asal inputnya, dari form atau file
        if (entries[0][1].get() == "Text"):
            result = entries[3][1].get()
        elif (entries[0][1].get() == "File Text"):
            text_file = open(filename, "r")
            result = text_file.read()
        # Hilangkan angka dan spasi
        if (entries[2][1].get() == "Extended Vigenere Cipher"):
            return result
        else:
            result = result.replace(" ", "")
            result = ''.join(i for i in result if not i.isdigit())
            result = result.translate(str.maketrans('', '', string.punctuation))
            return result.lower()
            
    elif (entries[0][1].get() == "File Random"):
        if (entries[1][1].get() == "Encrypt"):
            file = open(filename, "rb")
            f = file.read()
            b = bytearray(f)
            result = b.decode('latin-1')
            return(result)
        elif (entries[1][1].get() == "Decrypt"):
            file = open(filename, "rb")
            f = file.read()
            b = bytearray(f)
            result = b.decode('utf-8')
            return(result)
        

def getKey(entries):
    if (entries[2][1].get() == "Playfair Cipher"):
        return  entries[4][1].get()
    else:        
        result = []
        for elmt in entries[4][1].get().split(' ') :
            result.append(elmt.split(','))
        return result

def UploadAction(event=None):
    global filename
    filename = filedialog.askopenfilename()

#Milih menu
def chooseMenu(menu_label):
    if (menu_label == "Pilih Input"):
        return Pilih_Input
    elif (menu_label == "Pilih Aksi"):
        return Pilih_Aksi
    else:
        return Pilih_Cipher

#Buat menu
def makeMenu(root, fields, entries):
    for field in fields:
        variable = tk.StringVar(root)
        variable.set(field)
        opt = tk.OptionMenu(root, variable, *chooseMenu(field))
        opt.config(width=60)
        opt.pack()
        entries.append((field, variable))
    return entries

#Buat form
def makeForm(root, fields, entries):
    for field in fields:
        row = tk.Frame(root)
        lab = tk.Label(row, width=15, text=field, anchor='w')
        ent = tk.Entry(row, width=45)
        row.pack(side=tk.TOP, fill=tk.X, padx=5, pady=5)
        lab.pack(side=tk.LEFT)
        ent.pack(side=tk.RIGHT, expand=tk.YES, fill=tk.X)
        entries.append((field, ent))
    return entries

#Hasil dengan spasi 5 huruf
def getSpacedResult(ciphertext):
    result = ""
    for i in range(len(ciphertext)):
        if (i % 5 == 0):
            result += " "
        result += ciphertext[i]
    return(result)

if __name__ == '__main__':
    root = tk.Tk()
    root.title("Playfair, Extended Vigenere & Hill Cipher")

    entries = []
    ents = makeMenu(root, fields_menu, entries)
    root.bind('<Return>', (lambda event, e=ents: fetch(e)))
    
    ents = makeForm(root, fields_text, ents)
    root.bind('<Return>', (lambda event, e=ents: fetch(e)))

    result = StringVar()
    spaced_result = StringVar()

    Label(root, textvariable=result).pack()
    Label(root, textvariable=spaced_result).pack()

    b1 = tk.Button(root, text='Show',
                  command=(lambda e=ents: fetch(e)))
    b1.pack(side=tk.LEFT, padx=3, pady=5)
    b2 = tk.Button(root, text='Upload File', command=UploadAction)
    b2.pack(side=tk.LEFT, padx=3, pady=5)
    b2 = tk.Button(root, text='Exit', command=root.quit)
    b2.pack(side=tk.LEFT, padx=3, pady=5)

    root.mainloop()