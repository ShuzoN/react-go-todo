
import React from "react";
import { MaterialUiPickersDate } from "@material-ui/pickers/typings/date";
import { Todo } from "./Tasks";
import { DatePickerForm } from "./utils/DatePickerForm";

export const TodoDatePickerForm = (props: {
    todo: Todo,
    onChange: (todo: Todo) => void,
}): JSX.Element => {
    const handleDateChange = (date: MaterialUiPickersDate) => {
        const todo = { ...props.todo, deadLine: date };
        props.onChange(todo);
    };

    return (
        <DatePickerForm
            selectedDate={props.todo.deadLine}
            handleDateChange={handleDateChange}
        />
    );
};