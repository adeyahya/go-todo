import axios from "axios";
import { useQuery } from "@tanstack/react-query";
import { PaginatedTodo } from "@/schema/todo";

const fetchTodoList = async () => {
  const { data } = await axios.get<PaginatedTodo>("/todo");
  return data;
};

export const useTodoListQuery = () => {
  const res = useQuery({ queryKey: ["todos"], queryFn: fetchTodoList });
  return res;
};
