/**
 * Created by justin on 6/28/16.
 */

import { IFetchStrategy } from './IFetchStrategy';
import { IFetchStrategyConfig } from './IFetchStrategyConfig';

export class AJAXFetch implements IFetchStrategy {
  request(dispatch: Redux.Dispatch, config: IFetchStrategyConfig): void {
    return fetch(config.url)
      .then(response => response.json())
      .then(json => {
        return dispatch(config.callback(json));
      });
  }
}