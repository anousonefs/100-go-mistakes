package main

func main() {
	/* donationV1() */
	/* donationV2() */
	donationV3()
}

/*
chan struct vs sync.Cond:
- chan struct if there’s no active receiver, the message is buffered,
  which guarantees that this notification will be received eventually.
- Using sync.Cond with the Broadcast method wakes all goroutines currently waiting on the condition;
if there are none, the notification will be missed. This is also an essential principle that we have to keep in mind
*/

/*
sync.Cond.Wait():
	1. Unlock the mutex.
	2. Suspend the goroutine, and wait for a notification.
	3. Lock the mutex when the notification arrives.
*/
