
import React from "react";
import { MaterialUiPickersDate } from "@material-ui/pickers/typings/date";
import { DatePickerForm } from "./utils/DatePickerForm";
import { Todo } from "../Contract";

export const TodoDatePickerForm = (props: {
    todo: Todo,
    onChange: (todo: Todo) => void,
}): JSX.Element => {
    const handleDateChange = (date: MaterialUiPickersDate) => {
        const todo = { ...props.todo, deadline: date };
        props.onChange(todo);
    };

    return (
        <DatePickerForm
            selectedDate={props.todo.deadline}
            handleDateChange={handleDateChange}
        />
    );
};