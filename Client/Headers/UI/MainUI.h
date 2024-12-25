#ifndef ADAPTIXCLIENT_MAINUI_H
#define ADAPTIXCLIENT_MAINUI_H

#include <main.h>
#include <UI/Widgets/AdaptixWidget.h>
#include <Client/AuthProfile.h>

class MainUI : public QMainWindow
{
    QWidget* mainWidget = nullptr;

public:
    explicit MainUI();
    ~MainUI();

    void onExtender();

    void addNewProject(AuthProfile* profile);
    void addNewExtension(ExtensionFile extFile);
};

#endif //ADAPTIXCLIENT_MAINUI_H
