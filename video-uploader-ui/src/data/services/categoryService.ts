import { Category } from "../models/Category";
import http from "./../http.common";
const getAll = () => {
  return http.get<Array<Category>>("/categories");
};

const CategoryService = {
  getAll,
};

export default CategoryService;

