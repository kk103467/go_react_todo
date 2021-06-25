import React, { useState } from 'react'

export default function TodoForm(props) {
    const [taskname, setTaskname] = useState("");

    const handleSubmit = e => {
        e.preventDefault();
        if (!taskname) return;
        const textobj = {text: taskname}
        const textjson = JSON.stringify(textobj)
        props.addTodo(textjson);
        setTaskname("");
    };

    return (
        <form>
            <input
                type="text"
                value={taskname}
                onChange={e => { setTaskname(e.target.value) }}
            />
            <button
                type="submit"
                onClick={handleSubmit}
            >
                Add new task
            </button>
        </form>
    );
}