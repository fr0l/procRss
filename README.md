# procRss
A utility to get RSS for a process and descendants

```bash
./procRss 1234
```

```bash
pgrep -f java | xargs ./procRss

# sample output
[./procRss_linux64 5789 12889]
PID:    5789, process RSS      70 MB, descendants RSS    1587 MB, total RSS    1657 MB
PID:   12889, process RSS      74 MB, descendants RSS     685 MB, total RSS     759 MB
```