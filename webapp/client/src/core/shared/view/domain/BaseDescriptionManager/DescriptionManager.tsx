import { bind } from 'decko';
import * as React from 'react';
import { connect } from 'react-redux';
import { Dispatch, bindActionCreators } from 'redux';

import { EntityWithDescription } from 'core/shared/models/Description';
import { Icon } from 'core/shared/view/elements/Icon/Icon';
import Tooltip from 'core/shared/view/elements/Tooltip/Tooltip';
import { addOrEditDescription } from 'store/descriptionAction';

import AddEditDescModal from './AddEditDescModal/AddEditDescModal';
import styles from './DescriptionManager.module.css';

interface ILocalProps {
  entityId: string;
  entityType: EntityWithDescription;
  isReadOnly: boolean;
  isLoadingAccess: boolean;
  description: string;
}

interface IActionProps {
  addOrEditDescription: typeof addOrEditDescription;
}

interface ILocalState {
  isShowModal: boolean;
  showTooltip: boolean;
}

type AllProps = ILocalProps & IActionProps;

class DescriptionManager extends React.Component<AllProps, ILocalState> {
  public state: ILocalState = {
    isShowModal: false,
    showTooltip: false,
  };

  public render() {
    const { description, isLoadingAccess, isReadOnly } = this.props;
    return (
      <span className={styles.root} data-test="description">
        <span
          className={
            description ? styles.description : styles.description_empty
          }
          data-test="description-text"
        >
          {(() => {
            if (description) {
              return description;
            }
            if (!isLoadingAccess && !isReadOnly) {
              return 'Description';
            }
            return '';
          })()}
          {!isLoadingAccess && !isReadOnly && (
            <span>
              <div
                className={styles.desc_action_block}
                onClick={this.showModal}
                data-test="description-open-editor-button"
              >
                <Icon className={styles.desc_action_icon} type={'pencil'} />
              </div>
              <AddEditDescModal
                description={description}
                isOpen={this.state.isShowModal}
                onAddEdit={this.addEditDesc}
                onClose={this.closeModal}
              />
            </span>
          )}
          {!isLoadingAccess && isReadOnly && (
            <span>
              <Tooltip
                content={'Read Only Project'}
                visible={this.state.showTooltip}
              >
                <div
                  className={styles.desc_readonly}
                  onMouseOver={this.showTooltip}
                  onMouseOut={this.hideTooltip}
                >
                  <Icon className={styles.desc_readonly_icon} type={'eye'} />
                </div>
              </Tooltip>
            </span>
          )}
        </span>
      </span>
    );
  }

  @bind
  private showTooltip(event: React.MouseEvent) {
    event.preventDefault();
    this.setState({ showTooltip: true });
  }

  @bind
  private hideTooltip(event: React.MouseEvent) {
    event.preventDefault();
    this.setState({ showTooltip: false });
  }

  @bind
  private showModal(event: React.MouseEvent<HTMLDivElement, MouseEvent>) {
    event.preventDefault();
    this.setState({ isShowModal: true });
  }

  @bind
  private closeModal() {
    this.setState({ isShowModal: false });
  }

  @bind
  private addEditDesc(descCreated: string) {
    const { entityId: id, entityType } = this.props;
    this.props.addOrEditDescription(id, descCreated, entityType);
    this.setState({ isShowModal: false });
  }
}

const mapDispatchToProps = (dispatch: Dispatch): IActionProps =>
  bindActionCreators(
    {
      addOrEditDescription,
    },
    dispatch
  );

export type IDescriptionManagerLocalProps = ILocalProps;
export default connect(
  null,
  mapDispatchToProps
)(DescriptionManager);
