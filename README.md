# Logfile Analyzer

Retrieve, filter and review logs

## Retrieval of logfile

Using `scp`and assuming that the config file was set, replace the HOSTNAME and PATH in the following command and run it to copy the file:
```bash
scp HOSTNAME:PATH .
```

Since the logfiles have a prefix of 81 on every line that is not necessary for us, we can remove it using the following command:
```bash
sed 's/^.\{81\}//' OLD_FILE > NEW_FILE
```
To extract every log from the file following the pattern of [YYYY-MM-DD HH:MM:SS] we can use the following grep command to reduce the size further:
```bash
grep '^\[[0-9]\{4\}-[0-9]\{2\}-[0-9]\{2\} [0-9]\{2\}:[0-9]\{2\}:[0-9]\{2\}\]' FILENAME
```

