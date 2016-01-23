//
//  AppDelegate.swift
//  ComputerRoommate
//
//  Created by Dan Kurtz on 1/22/16.
//  Copyright © 2016 Dan Kurtz. All rights reserved.
//

import Cocoa

@NSApplicationMain
class AppDelegate: NSObject, NSApplicationDelegate {

    @IBOutlet weak var statusMenu: NSMenu?
    var statusItem: NSStatusItem? = nil;
    
    @IBAction func statusMenuItemSelected(sender: AnyObject) {
    
    }
    
    func applicationDidFinishLaunching(aNotification: NSNotification) {
        // Insert code here to initialize your application
    }

    func applicationWillTerminate(aNotification: NSNotification) {
        // Insert code here to tear down your application
    }
    
    override func awakeFromNib() {
        activateStatusMenu()
    }
    
    func activateStatusMenu() {
        let statusBar = NSStatusBar.systemStatusBar()
        
        statusItem = statusBar.statusItemWithLength(NSVariableStatusItemLength)
        statusItem!.title = "👩"
        statusItem!.menu = self.statusMenu
        NSApplication.sharedApplication().keyWindow?.backgroundColor = NSColor.whiteColor()
    }


}

