import { Moment } from "moment";

export interface Todo {
  id: number;
  title: string;
  deadline: Moment | null;
}
