# Vigenère Cipher
* Best known and one of the simplest polyalphabetic substitution ciphers
* In this scheme the set of related monoalphabetic substitution rules consists of the 26 Caesar ciphers with shifts of 0 through 25
* Each cipher is denoted by a key letter which is the ciphertext letter that substitutes for the plaintext letter 
## Example of Vigenère Cipher
![image](https://user-images.githubusercontent.com/62750835/118081369-e3656600-b3c3-11eb-9e69-6c1d4afa2daf.png)
* To encrypt a message, a key is needed that is as long as the message
 * Usually, the key is a repeating keyword 
* For example, if the keyword is deceptive, the message “we are discovered save yourself” is encrypted as:
  * key:          deceptivedeceptivedeceptive
  * plaintext:    wearediscoveredsaveyourself
  * ciphertext:  ZICVTWQNGRZGVTWAVZHCQYGLMGJ
