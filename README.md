Keeping attack surface minimal is one of the security best practices. Rogueport
helps you identify TCP ports which are not supposed to be open. It scans the
hosts defined in the config file for open ports. Then it compares the scan
results with the expected state.

Install binary:

```
git clone git@github.com:jreisinger/rogueport.git
cd rogueport
go install
```

Define ports you need to have open (i.e. you're running services on them), for
example:

```
$ cat rogueport.json
[
    {
        "host": "scanme.nmap.org",
        "ports": [ 22 ]
    },
    {
        "host": "scanme2.nmap.org",
        "ports": [ 22, 80, 443 ]
    }
]
```

Check there are no unexpected ports open:

```
$ rogueport
scanme.nmap.org           22 ✓ 80 ✗
scanme2.nmap.org          22 ✓ 25 ✗ 80 ✓ 443 ✓
```