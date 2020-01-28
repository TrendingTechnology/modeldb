import cn from 'classnames';
import { bind } from 'decko';
import * as R from 'ramda';
import * as React from 'react';
import { connect } from 'react-redux';
import { RouteComponentProps, withRouter } from 'react-router-dom';

import {
  IFilterData,
  IQuickFilter,
  PropertyType,
  IIdFilterData,
} from 'core/shared/models/Filters';
import Button from 'core/shared/view/elements/Button/Button';
import Droppable from 'core/shared/view/elements/Droppable/Droppable';
import IdView from 'core/shared/view/elements/IdView/IdView';
import {
  IFilterContext,
  IFilterContextData,
  resetCurrentContext,
  setContext,
  selectCurrentContextData,
  removeFilterFromCurrentContext,
  addFilterToCurrentContext,
  editFilterInCurrentContext,
  saveEditedFilterInCurrentContext,
} from 'store/filter';
import { IApplicationState, IConnectedReduxProps } from 'store/store';

import styles from './FilterManager.module.css';
import InstantFilterItem from './InstantFilterItem/InstantFilterItem';
import RemoveButton from './InstantFilterItem/shared/RemoveButton/RemoveButton';
import QuickFilterInput from './QuickFilterInput/QuickFilterInput';

interface ILocalProps {
  placeholderText: string;
  isCollapsed: boolean;
  context: IFilterContext;
  isDisabled: boolean;
  withFilterIdsSection?: boolean;
  onExpandSidebar(): void;
}

interface IPropsFromState {
  filters: IFilterData[];
  ctxName?: string;
  quickFilters: IQuickFilter[];
}

type AllProps = ILocalProps &
  IPropsFromState &
  IConnectedReduxProps &
  RouteComponentProps;

class FilterManager extends React.Component<AllProps> {
  public componentDidMount() {
    this.props.dispatch(setContext(this.props.context));
  }

  public componentWillUnmount() {
    this.props.dispatch(resetCurrentContext());
  }

  public render() {
    const {
      isCollapsed,
      filters,
      placeholderText,
      withFilterIdsSection,
      isDisabled,
      onExpandSidebar,
    } = this.props;
    const idFilter = (filters.find(
      filter => filter.type === PropertyType.ID
    ) as IIdFilterData) || {
      id: '',
      type: PropertyType.ID,
      name: 'id',
      value: [],
      isEdited: false,
    };
    const instantFilters = filters.filter(
      filter => filter.type !== PropertyType.ID
    );
    return (
      <div
        className={cn(styles.root, {
          [styles.collapsed]: isCollapsed,
          [styles.disabled]: isDisabled,
        })}
        data-test="filters"
      >
        <div className={cn(styles.quick_filter_input)}>
          <QuickFilterInput
            isCollapsed={isCollapsed}
            quickFilters={this.props.quickFilters}
            onCreateFilter={this.onCreateFilter}
            onExpandSidebar={onExpandSidebar}
          />
        </div>
        <div className={styles.filters}>
          <Droppable type="filter" onDrop={this.onCreateFilter}>
            <div
              className={cn(styles.instant_filters, {
                [styles.instant_filters_empty]: instantFilters.length === 0,
              })}
              data-test="filter-items-area"
            >
              {instantFilters.length === 0 ? (
                <div className={styles.instant_filters_placeholder}>
                  {placeholderText}
                </div>
              ) : (
                instantFilters
                  .filter(filter => filter.type !== PropertyType.ID)
                  .map((filter, index) => (
                    <InstantFilterItem
                      key={index}
                      data={filter}
                      onRemoveFilter={this.onRemoveFilter}
                      onChange={this.makeOnEditFilterData}
                    />
                  ))
              )}
            </div>
          </Droppable>
          {withFilterIdsSection && (
            <div className={styles.ids_filter_section}>
              <div
                className={cn(styles.ids_filter, {
                  [styles.ids_filter_empty]: idFilter.value.length === 0,
                })}
              >
                {idFilter && idFilter.value.length !== 0 ? (
                  idFilter.value.map(id => (
                    <div className={styles.ids_filter_item} key={id}>
                      <div
                        className={styles.ids_filter_item_remove}
                        data-test="id-filter-value"
                      >
                        <RemoveButton
                          onRemove={this.makeOnRemoveIdFilterItem(id)}
                        />
                      </div>
                      <div className={styles.ids_filter_item_name} title={id}>
                        <IdView value={id} />
                      </div>
                    </div>
                  ))
                ) : (
                  <div className={styles.id_filters_placeholder}>
                    Add to collection
                  </div>
                )}
              </div>
              <div className={styles.id_action_section}>
                <div className={styles.id_action_button}>
                  <Button
                    fullWidth={true}
                    dataTest="delete-ids-filter-button"
                    onClick={this.onClearFilter}
                  >
                    Clear
                  </Button>
                </div>
                <div className={styles.id_action_button}>
                  <Button
                    fullWidth={true}
                    dataTest="apply-ids-filter-button"
                    onClick={this.onSaveIdFilter}
                  >
                    Filter
                  </Button>
                </div>
              </div>
            </div>
          )}
        </div>
      </div>
    );
  }

  @bind
  private onCreateFilter(data: IFilterData) {
    const { dispatch } = this.props;
    dispatch(addFilterToCurrentContext(data));
  }

  @bind
  private onRemoveFilter(data: IFilterData) {
    this.props.dispatch(removeFilterFromCurrentContext(data));
  }

  @bind
  private makeOnEditFilterData(data: IFilterData) {
    this.props.dispatch(editFilterInCurrentContext(data));
  }

  @bind
  private onSaveIdFilter() {
    const idFilter = this.props.filters.find(
      filter => filter.type === PropertyType.ID
    ) as IIdFilterData;
    if (idFilter) {
      this.props.dispatch(saveEditedFilterInCurrentContext(idFilter));
    }
  }
  @bind
  private onClearFilter() {
    const idFilter = this.props.filters.find(
      filter => filter.type === PropertyType.ID
    ) as IIdFilterData;
    if (idFilter) {
      this.props.dispatch(removeFilterFromCurrentContext(idFilter));
    }
  }
  @bind
  private makeOnRemoveIdFilterItem(modelId: string) {
    return () => {
      const idFilter = this.props.filters.find(
        filter => filter.type === PropertyType.ID
      ) as IIdFilterData;
      if (idFilter) {
        const updatedIdFilter: IIdFilterData = {
          ...idFilter,
          isEdited: true,
          value: R.without([modelId], idFilter.value),
        };
        this.props.dispatch(editFilterInCurrentContext(updatedIdFilter));
      }
    };
  }
}

const mapStateToProps = (state: IApplicationState): IPropsFromState => {
  const currentContextData:
    | IFilterContextData
    | undefined = selectCurrentContextData(state);
  const quickFilters = currentContextData
    ? currentContextData.ctx.quickFilters
    : [];

  return {
    filters: currentContextData ? currentContextData.filters : [],
    quickFilters,
  };
};

export type IFilterManagerLocalProps = ILocalProps;
export default withRouter(connect(mapStateToProps)(FilterManager));
