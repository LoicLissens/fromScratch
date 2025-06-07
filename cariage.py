import time, sys
total = 10
print('Progress:')
for i in range(total):
    sys.stdout.write('%d / %d\r' % (i, total+i))
    sys.stdout.flush()
    time.sleep(0.5)
print('Done     ')