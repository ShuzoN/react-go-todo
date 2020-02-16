import { TodoGateway } from "./Gateway/TodoGateway";
import { Todo } from "./Contract";

export async function fetchTodos(
  todoGateway: TodoGateway
): Promise<Todo[] | Error> {
  const res = await todoGateway.getAll();
  if (res.status >= 400) {
    throw new Error(res.status + ": fetch fail");
  }
  const jsonObj: Promise<Todo[]> = await res.json();
  const todos: Todo[] = (await jsonObj).map(todo => {
    return { ...todo, deadline: todo.deadline };
  });
  return todos;
}

export async function updateTodo(
  todoGateway: TodoGateway,
  todo: Todo
): Promise<Todo | Error> {
  const res = await todoGateway.update(todo);
  if (res.status >= 400) {
    throw new Error(res.status + ": update fail");
  }
  return todo;
}

export async function createTodo(
  todoGateway: TodoGateway,
  todo: Todo
): Promise<Todo | Error> {
  const res = await todoGateway.create(todo);
  if (res.status >= 400) {
    throw new Error(res.status + ": update fail");
  }
  return todo;
}
