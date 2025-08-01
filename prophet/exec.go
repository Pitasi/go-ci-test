package prophet

import (
	"context"
	"log/slog"
)

// TaskResultWriter writes results of task executions.
type TaskResultWriter interface {
	Write(result TaskResult) error
}

// ExecTasks executes tasks coming from the specified reader and writes
// them to the specified writer.
//
// This call is non-blocking, the main loop is executed in a goroutine.
func ExecTasks(r TaskReader, w TaskResultWriter) error {
	log := slog.With("process", "exec_tasks")

	go func() {
		for task := range r.Read() {
			log := log.With("task", task.ID)
			ctx := context.TODO()

			log.DebugContext(ctx, "running task")

			output, err := Execute(ctx, task)
			if err != nil {
				log.ErrorContext(ctx, "failed to run task", "err", err)
				continue
			}

			err = w.Write(output)
			if err != nil {
				log.ErrorContext(ctx, "failed to write task result", "err", err)
				continue
			}
		}
	}()

	return nil
}

// VoteWriter writes votes of task verifications.
type VoteWriter interface {
	Write(result Vote) error
}

// ExecVotes executes task verifications coming from the specified reader and
// writes them to the specified writer.
//
// This call is non-blocking, the main loop is executed in a goroutine.
func ExecVotes(r TaskResultReader, w VoteWriter) error {
	log := slog.With("process", "exec_votes")

	go func() {
		for proposal := range r.Read() {
			plog := log.With("task", proposal.ID)
			ctx := context.TODO()

			plog.DebugContext(ctx, "verifying task proposal")

			err := Verify(ctx, proposal)
			if err := w.Write(Vote{
				ID:  proposal.ID,
				Err: err,
			}); err != nil {
				plog.ErrorContext(ctx, "failed to write vote", "err", err)
				continue
			}
		}
	}()

	return nil
}
