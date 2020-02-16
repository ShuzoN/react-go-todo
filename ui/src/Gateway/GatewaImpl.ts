import { Gateway } from "./Gateway";

const ORIGIN = "http://localhost:80";

export class GatewayImpl implements Gateway {
  async get(path: string): Promise<Response> {
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
        return res;
      })
      .catch(rejected => {
        throw new Error(rejected);
      });
  }
  async post(path: string, json: string): Promise<Response> {
    console.log(path);
    console.log(json);
    throw new Error("Method not implemented.");
  }
}
