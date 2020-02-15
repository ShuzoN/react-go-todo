import { Gateway } from "./Gateway";
export interface TodoGateway {
  getById(id: number): Promise<Response>;
  getAll(): Promise<Response>;
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
}
