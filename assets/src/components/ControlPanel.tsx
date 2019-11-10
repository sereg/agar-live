import * as React from 'react';
import {Status} from '../const/Const';
import {Play, Stop} from './icon/icon';

interface Props {
    status: Status;
    changes: (event: React.MouseEvent<HTMLDivElement>) => void
}

export const ControlPanel = (props: Props) => {

    return (
        <div className="ff" onClick={props.changes}>
            {props.status === Status.playing ? (
                <Play width="20px"/>
            ) : (
                <Stop width="20px"/>
            )}
        </div>
    )
}