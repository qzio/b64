# b64

Mini version of `base64` but using url safe base64 encoding with padding stripped by default. (as per jwt standard)
To disable stripping, use the `-nostrip` flag.

```
# encode
echo 'foo' | b64

Zm9vCg==

# decode
echo 'Zm9vCg==' | b64 -d

foo

```

