.DEFAULT_GOAL:= exec

.PHONY: version
version:
	$(MAKE) -C show-version

.PHONY: buffering
buffering:
	$(MAKE) -C buffering

.PHONY: multi-writer
multi-writer:
	$(MAKE) -C multi-writer

.PHONY: rand
rand:
	$(MAKE) -C rand

.PHONY: humanize
humanize:
	$(MAKE) -C humanize

.PHONY: exec
exec:
	$(MAKE) -C exec

.PHONY: no-signal
no-signal:
	$(MAKE) -C no-signal

.PHONY: handle-signal
handle-signal:
	$(MAKE) -C handle-signal

.PHONY: timer-signal
timer-signal:
	$(MAKE) -C timer-signal

.PHONY: stop-goroutine
stop-goroutine:
	$(MAKE) -C stop-goroutin-by-channel
	$(MAKE) -C stop-goroutin-by-context
