import { Gateway } from "./Gateway";
export interface UserGateway {
  getById(id: number): Promise<Response>;
}

export class UserGatewayImpl implements UserGateway {
  private gateway: Gateway;
  constructor(gateway: Gateway) {
    this.gateway = gateway;
  }

  getById(id: number): Promise<Response> {
    return this.gateway.get("/users/" + id);
  }
}
