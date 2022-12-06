# tiny wait-until tool

If you want pause a script or delay a program execution until a specific time.

```bash
wu h[:m]
```

Example:

```bash
wu 10 && make build
```

If the time passed today, it waits until tomorrow. It supports also negative numbers.

Let's say it's 9:15 (am):

```bash
wu 9               # wait until tomorrow 9:00
wu 10              # wait until 10:00
wu 16              # wait until 16:00
wu 24              # wait until tomorrow 0:00
wu 25              # wait until tomorrow 1:00
wu -2              # wait until 22:00

wu 9:10            # wait until tomorrow 9:10
wu 9:20            # wait until 9:20
wu 16:30           # wait until 16:30
wu 26:90           # wait until tomorrow 3:30
wu 24:-5           # wait until 23:55
```
