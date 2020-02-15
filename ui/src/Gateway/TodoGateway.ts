const ORIGIN = 'http://localhost:80/';

export interface TodoGateway {
  get(path: string): Promise<Response>;
  post(path: string, json: JSON): Promise<Response>;
}

export class TodoGatewayImpl implements TodoGateway {
  get(path: string): Promise<Response> {
    return this._get<Response>(path);
  }
  async _get<T>(path: string): Promise<T> {
    return await fetch(ORIGIN + , { mode: "cors" }).then(res => {
      if (!res.ok) {
        throw new Error(res.statusText);
      }
      return res.json();
    });
  }
  post(path: string, json: JSON): Promise<Response> {
    throw new Error("Method not implemented.");
  }
  async _post<T>(path: string, json: JSON): Promise<T> {
    throw new Error("Method not implemented.");
  }
}
