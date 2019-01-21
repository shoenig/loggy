# loggy

Go's worst logging package.

### Usage
```
log := New("my-thing")
log.Tracef("this is trace %d", 1)
log.Infof("this is info %d", 2)
log.Warnf("this is warn %d", 3)
log.Errorf("this is error %d", 4)
```

Output:
```
2019/01/21 09:36:32 TRACE [my-thing] this is trace 1
2019/01/21 09:36:32 INFO  [my-thing] this is info 2
2019/01/21 09:36:32 WARN  [my-thing] this is warn 3
2019/01/21 09:36:32 ERROR [my-thing] this is error 4
```
