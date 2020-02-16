export interface Gateway {
  get(path: string): Promise<Response>;
  post(path: string, json: string): Promise<Response>;
}
