import { PaginatedTodo, Todo } from "@/schema/todo";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import axios from "axios";

type Payload = {
  id: string;
  title: string;
};

export const useTodoCreateMutation = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: async (payload: Payload) => {
      const { data } = await axios.post<Todo>("/todo", payload);
      return data;
    },
    onMutate: (payload: Payload) => {
      const optimisticData: Todo = {
        ...payload,
        isCompleted: false,
        createdAt: new Date().toISOString(),
        completedAt: null,
      };

      queryClient.setQueryData(["todos"], (old: PaginatedTodo | undefined) => {
        if (!old) return { cursor: "", data: [optimisticData] };

        const data = old.data;
        old.data = [optimisticData, ...data];
        return old;
      });

      return { optimisticData };
    },
    onError: (_, __, context) => {
      if (context?.optimisticData) {
        queryClient.setQueryData(
          ["todos"],
          (old: PaginatedTodo | undefined) => {
            if (!old) return old;
            return {
              ...old,
              data: old.data.filter(
                (item) => item.id !== context.optimisticData.id
              ),
            };
          }
        );
      }
    },
  });
};
