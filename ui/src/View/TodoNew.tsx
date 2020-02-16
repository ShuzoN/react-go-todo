import React, { useState } from 'react';
import { gateways } from '../App';
import { Todo, initTodo } from '../Contract';
import { TodoEditView } from './TodoEdit';
import { createTodo } from '../Epic';

export const TodoNew = (props: {
}): JSX.Element => {
    const [todo, setTodo] = useState<Todo>(initTodo());

    return <TodoEditView
        todo={todo}
        onChange={
            (todo) => {
                setTodo(todo);
                createTodo(gateways.todoGateway, todo);
            }
        }
    />
};
