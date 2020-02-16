import { Moment } from "moment";

export interface Todo {
  id: number;
  title: string;
  deadline: Moment | null;
  checked: boolean;
}

export function initTodo(): Todo {
  return {
    id: 0,
    title: "",
    deadline: null,
    checked: false
  };
}
