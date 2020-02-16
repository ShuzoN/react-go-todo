import { Gateway } from "./Gateway";
import { Todo } from "../Contract";
import moment from "moment";

export interface TodoGateway {
  getById(id: number): Promise<Response>;
  getAll(): Promise<Response>;
  update(todo: Todo): Promise<Response>;
  create(todo: Todo): Promise<Response>;
}

export class TodoGatewayImpl implements TodoGateway {
  private gateway: Gateway;
  constructor(gateway: Gateway) {
    this.gateway = gateway;
  }

  getById(id: number): Promise<Response> {
    return this.gateway.get("/todos/" + id);
  }

  getAll(): Promise<Response> {
    return this.gateway.get("/todos/");
  }

  create(todo: Todo): Promise<Response> {
    return this.gateway.post("/todos/", JSON.stringify(todo, todoSerializer));
  }

  update(todo: Todo): Promise<Response> {
    return this.gateway.post(
      "/todos/" + todo.id,
      JSON.stringify(todo, todoSerializer)
    );
  }
}

function todoSerializer(key: any, value: any) {
  if (key === "deadline") {
    return moment(value).format("YYYY-MM-DDTHH:mm:ssZ");
  }

  return value;
}
