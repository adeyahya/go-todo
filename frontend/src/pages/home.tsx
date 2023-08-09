import {
  Container,
  Input,
  Stack,
  Text,
  Box,
  HStack,
  Checkbox,
} from "@chakra-ui/react";
import { useTodoCreateMutation, useTodoListQuery } from "@/hooks";
import { useForm } from "react-hook-form";
import { generateId } from "@/utils";

const Home = () => {
  const { handleSubmit, register, resetField } = useForm<{ title: string }>();
  const { data } = useTodoListQuery();
  const createTodo = useTodoCreateMutation();
  const { data: todoList = [] } = data ?? {};

  const onSubmit = handleSubmit(async (data) => {
    const id = generateId();
    await createTodo.mutateAsync({ id, title: data.title });
    resetField("title");
  });

  return (
    <Container>
      <Stack>
        <Box mt="4" as="form" onSubmit={onSubmit}>
          <Input placeholder="New Todo" {...register("title")} />
        </Box>
        {todoList.map((todo) => (
          <HStack key={todo.id}>
            <Box>
              <Checkbox size="lg" isChecked={todo.isCompleted} />
            </Box>
            <Text>{todo.title}</Text>
          </HStack>
        ))}
      </Stack>
    </Container>
  );
};

export default Home;
