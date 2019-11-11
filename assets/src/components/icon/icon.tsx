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

export const Refresh = (props: Props) => {

    let width: string = props.width || "20px";

    return (
        <svg style={{width: width}} viewBox="0 0 512 512">
            <path d="M512 192h-192l71.765-71.765c-36.265-36.263-84.48-56.235-135.765-56.235s-99.5 19.972-135.765 56.235c-36.263 36.265-56.235 84.48-56.235 135.765s19.972 99.5 56.235 135.765c36.265 36.263 84.48 56.235 135.765 56.235s99.5-19.972 135.764-56.236c3.028-3.027 5.93-6.146 8.728-9.334l48.16 42.141c-46.923 53.583-115.832 87.429-192.652 87.429-141.385 0-256-114.615-256-256s114.615-256 256-256c70.693 0 134.684 28.663 181.008 74.992l74.992-74.992v192z"></path>
        </svg>
    )
}

export const Backward = (props: Props) => {

    let width: string = props.width || "20px";

    return (
        <svg style={{width: width}} viewBox="0 0 512 512">
            <path d="M288 80v160l160-160v352l-160-160v160l-176-176z"></path>
        </svg>
    )
}

export const Forward = (props: Props) => {

    let width: string = props.width || "20px";

    return (
        <svg style={{width: width}} viewBox="0 0 512 512">
            <path d="M256 432v-160l-160 160v-352l160 160v-160l176 176z"></path>
        </svg>
    )
}

export const Save = (props: Props) => {

    let width: string = props.width || "20px";

    return (
        <svg style={{width: width}} viewBox="0 0 512 512">
            <path d="M224 288h64v-128h96l-128-128-128 128h96zM320 216v49.356l146.533 54.644-210.533 78.509-210.533-78.509 146.533-54.644v-49.356l-192 72v128l256 96 256-96v-128z"></path>
        </svg>
    )
}

export const Load = (props: Props) => {

    let width: string = props.width || "20px";

    return (
        <svg style={{width: width}} viewBox="0 0 512 512">
            <path d="M256 288l128-128h-96v-128h-64v128h-96zM372.363 235.636l-35.87 35.871 130.040 48.493-210.533 78.509-210.533-78.509 130.040-48.493-35.871-35.871-139.636 52.364v128l256 96 256-96v-128z"></path>
        </svg>
    )
}