import { Category } from "./Category";

export interface Video {
  id: number;
  title: string;
  categoryID: number;
  thumbnail256:string;
  Category: Category;
}