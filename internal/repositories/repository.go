package repositories

type Task interface{}

type Repository struct {
	Task *TaskRepository
}

func NewRepository() *Repository {
	return &Repository{
		Task: NewTaskRepository(),
	}
}
