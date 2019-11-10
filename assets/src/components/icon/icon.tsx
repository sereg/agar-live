import React from 'react';

interface Props {
    width: string
}

export const Play = (props: Props) => {

    let width: string = props.width || "20px";

    return (
        <svg style={{width: width}} viewBox="0 0 512 512">
            <path d="M96 64l320 192-320 192z"></path>
        </svg>
    )
}

export const Stop = (props: Props) => {

    let width: string = props.width || "20px";

    return (
        <svg style={{width: width}} viewBox="0 0 512 512">
            <path d="M64 64h160v384h-160zM288 64h160v384h-160z"></path>
        </svg>
    )
}