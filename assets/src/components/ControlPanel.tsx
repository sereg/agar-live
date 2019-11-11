import * as React from 'react';
import {Status} from '../const/Const';
import {Play, Stop} from './icon/icon';

interface Props {
    status: Status;
    changes: (event: React.MouseEvent<HTMLDivElement>) => void
    restart: (event: React.MouseEvent<HTMLDivElement>) => void
}

export const ControlPanel = (props: Props) => {

    return (
        <div>
            <div onClick={props.changes}>
                {props.status === Status.playing ? (
                    <Play width="50px"/>
                ) : (
                    <Stop width="50px"/>
                )}
            </div>
            <div onClick={props.restart}>
                <Play width="50px"/>
            </div>
        </div>
    )
}