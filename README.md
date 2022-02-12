Keeping attack surface minimal is one of the security best practices. Rogueport
helps you identify network ports which are not supposed to be open. It scans the
hostnames defined in the config file for open ports. Then it compares the scan
results with the expected state defined in the config file. NOTE: scan only your
hosts or hosts you have [permission](http://scanme.nmap.org/) to scan!

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
        "hostname": "scanme.nmap.org",
        "ports": [ "22/tcp" ]
    },
    {
        "hostname": "scanme2.nmap.org",
        "ports": [ "22/tcp", "80/tcp", "443/tcp" ]
    }
]
```

Check there are no unexpected ports open:

```
$ rogueport
scanme.nmap.org           22/tcp ✓ 80/tcp ✗
scanme2.nmap.org          22/tcp ✓ 25/tcp ✗ 80/tcp ✓ 443/tcp ✓
```

Rogueport uses [nmap](https://nmap.org/) to do the scanning, so you need to have
it installed (e.g. `apt get install nmap` or `brew install nmap`).
