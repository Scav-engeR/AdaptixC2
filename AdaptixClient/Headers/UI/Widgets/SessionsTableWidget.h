#ifndef ADAPTIXCLIENT_SESSIONSTABLEWIDGET_H
#define ADAPTIXCLIENT_SESSIONSTABLEWIDGET_H

#include <main.h>
#include <Agent/Agent.h>

class SessionsTableWidget : public QWidget
{
    QWidget*          mainWidget     = nullptr;
    QGridLayout*      mainGridLayout = nullptr;
    QTableWidget*     tableWidget    = nullptr;
    QMenu*            menuSessions   = nullptr;
    QTableWidgetItem* titleAgentID   = nullptr;
    QTableWidgetItem* titleAgentType = nullptr;
    QTableWidgetItem* titleListener  = nullptr;
    QTableWidgetItem* titleExternal  = nullptr;
    QTableWidgetItem* titleInternal  = nullptr;
    QTableWidgetItem* titleDomain    = nullptr;
    QTableWidgetItem* titleComputer  = nullptr;
    QTableWidgetItem* titleUser      = nullptr;
    QTableWidgetItem* titleOs        = nullptr;
    QTableWidgetItem* titleProcess   = nullptr;
    QTableWidgetItem* titleProcessId = nullptr;
    QTableWidgetItem* titleThreadId  = nullptr;
    QTableWidgetItem* titleTag       = nullptr;
    QTableWidgetItem* titleLast      = nullptr;
    QTableWidgetItem* titleSleep     = nullptr;

    void createUI();

public:
    int ColumnAgentID   = 0;
    int ColumnAgentType = 1;
    int ColumnExternal  = 2;
    int ColumnListener  = 3;
    int ColumnInternal  = 4;
    int ColumnDomain    = 5;
    int ColumnComputer  = 6;
    int ColumnUser      = 7;
    int ColumnOs        = 8;
    int ColumnProcess   = 9;
    int ColumnProcessId = 10;
    int ColumnThreadId  = 11;
    int ColumnTags      = 12;
    int ColumnLast      = 13;
    int ColumnSleep     = 14;
    int ColumnCount     = 15;

    explicit SessionsTableWidget( QWidget* w );
    ~SessionsTableWidget() override;

    void Clear() const;
    void AddAgentItem(Agent* newAgent) const;
    void RemoveAgentItem(const QString &agentId) const;

public slots:
    void handleTableDoubleClicked( const QModelIndex &index ) const;
    void handleSessionsTableMenu(const QPoint &pos );

    void actionConsoleOpen() const;
    void actionTasksBrowserOpen() const;
    void actionFileBrowserOpen() const;
    void actionProcessBrowserOpen() const;
    void actionAgentExit() const;
    void actionAgentTag() const;
    void actionAgentHide() const;
    void actionAgentRemove() const;
};

#endif //ADAPTIXCLIENT_SESSIONSTABLEWIDGET_H
