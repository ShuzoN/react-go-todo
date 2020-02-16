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
