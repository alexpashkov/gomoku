import React from "react";

interface INotificaitonsProps<T> {
  timeout: number;
  children: (notifications: T[], notify: (x: T) => void) => React.ReactNode;
}

interface INotificationsState<T> {
  notifications: { notification: T; created: number }[];
}

class Notifications<T> extends React.Component<
  INotificaitonsProps<T>,
  INotificationsState<T>
> {
  notify = (notification: T) =>
    this.setState(({ notifications }) => ({
      notifications: [
        {
          notification,
          created: Date.now()
        },
        ...notifications
      ]
    }));

  render() {
    return this.props.children(
      this.state.notifications.map(x => x.notification),
      this.notify
    );
  }
}

export default Notifications;
