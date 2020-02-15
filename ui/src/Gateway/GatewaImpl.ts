import { Gateway } from "./Gateway";

const ORIGIN = "http://localhost:80";

export class GatewayImpl implements Gateway {
  get(path: string): Promise<Response> {
    return this._get<Response>(path);
  }
  async _get<T>(path: string): Promise<T> {
    return await fetch(ORIGIN + path, {
      mode: "cors",
      method: "get",
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    })
      .then(res => {
        if (!res.ok) {
          throw new Error(res.statusText);
        }
        return res.json();
      })
      .catch(rejected => {
        throw new Error(rejected);
      });
  }
  post(path: string, json: JSON): Promise<Response> {
    throw new Error("Method not implemented.");
  }

  async _post<T>(path: string, json: JSON): Promise<T> {
    throw new Error("Method not implemented.");
  }
}