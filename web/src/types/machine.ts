import { Interface } from "./interface";

export interface Machine {
  name: string;
  id: string;
  groups: string[];
  node_id: string;
  hostname: string;
  interfaces: Interface[];
}
